package main

import (
	"fmt"
	"log"
	"log-converter/data"
	"log-converter/model"
	"log-converter/parser"
	"log-converter/reader"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func main() {
	fileMap := map[string][]model.File{
		model.FirstFormat:  []model.File{},
		model.SecondFormat: []model.File{},
	}

	for _, file := range getFiles() {
		if files, ok := fileMap[file.LogFormat]; ok {
			fileMap[file.LogFormat] = append(files, file)
		} else {
			log.Printf("Unsupported log format: %s.\n", file.LogFormat)
		}
	}

	c := make(chan model.Entry, 10)

	r := reader.New()

	for logFormat, files := range fileMap {
		if len(files) == 0 {
			continue
		}
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()

		for _, file := range files {
			err = watcher.Add(file.FilePath)
			if err != nil {
				log.Fatal(err)
			}
		}

		p := parser.New(parser.GetParseDate(logFormat), logFormat)

		go watch(watcher, r, p, c)
	}

	for e := range c {
		if _, err := data.CreateLogEntry(e); err != nil {
			log.Println(err)
		}
	}
}

func getFiles() []model.File {
	rawFiles := os.Args[1:]

	if len(rawFiles) == 0 {
		fmt.Println("Please provide file list.")
	}

	files := make([]model.File, 0)

	for _, rawFile := range rawFiles {
		rawMeta := strings.Split(rawFile, "|")

		if len(rawMeta) != 2 {
			fmt.Printf("Incorrect argument format: \"%s\". Expected: \"file path|log format\".\n", rawFile)
			continue
		}

		files = append(
			files,
			model.File{
				FilePath:  strings.Trim(rawMeta[0], " \t"),
				LogFormat: strings.Trim(rawMeta[1], " \t"),
			})
	}

	return files
}

func watch(w *fsnotify.Watcher, r *reader.Reader, p *parser.Parser, c chan<- model.Entry) {
	for {
		select {
		case event := <-w.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				strs := r.Read(event.Name)
				for _, str := range strs {
					str = strings.Trim(str, " \t\n\r")
					if len(str) == 0 {
						continue
					}
					entry, err := p.Parse(str, event.Name)
					if err != nil {
						continue
					}
					c <- entry
				}
			}
		case err := <-w.Errors:
			log.Println(err)
		}
	}
}
