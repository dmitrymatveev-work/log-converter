package main

import (
	"fmt"
	"log"
	"log-converter/model"
	"log-converter/reader"
	"os"
	"strings"
	"time"

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

		go watch(watcher, r, getParseDate(logFormat), logFormat, c)
	}

	for e := range c {
		store(e)
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
			fmt.Printf("Wrong argument format: \"%s\". Expected: \"file path|log format\".\n", rawFile)
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

func getParseDate(logFormat string) func(s string) time.Time {
	switch logFormat {
	case model.FirstFormat:
		return parseFirstDate
	case model.SecondFormat:
		return parseSecondDate
	}
	panic(fmt.Sprintf("Unsupported log format: %s.\n", logFormat))
}

func parseFirstDate(s string) time.Time {
	panic("the function is not implemented")
}

func parseSecondDate(s string) time.Time {
	panic("the function is not implemented")
}

func watch(w *fsnotify.Watcher, r *reader.Reader, parseDate func(string) time.Time, logFormat string, c chan<- model.Entry) {
	for {
		select {
		case event := <-w.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				strs := r.Read(event.Name)
				for _, str := range strs {
					fmt.Println(str)
				}
				log.Printf("Write to %s.\n", logFormat)
			}
		case err := <-w.Errors:
			log.Println(err)
		}
	}
}

func store(e model.Entry) {
	panic("the function is not implemented")
}
