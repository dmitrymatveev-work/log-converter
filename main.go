package main

import (
	"fmt"
	"log-converter/model"
	"os"
	"strings"
	"time"
)

func main() {
	c := make(chan model.Entry, 10)

	parseFirstDate := func(s string) time.Time {
		panic("the function is not implemented")
	}

	parseSecondDate := func(s string) time.Time {
		panic("the function is not implemented")
	}

	for _, file := range getFiles() {
		switch file.LogFormat {
		case "first_format":
			discover(file.FilePath, parseFirstDate, c)
		case "second_format":
			discover(file.FilePath, parseSecondDate, c)
		}
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

func discover(filePath string, parseDate func(string) time.Time, c chan<- model.Entry) {
	fmt.Printf("Discovering %s\n", filePath)
}

func store(e model.Entry) {
	panic("the function is not implemented")
}
