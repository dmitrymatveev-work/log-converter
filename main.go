package main

import "log-converter/model"

func main() {
	c := make(chan model.Entry, 10)

	for _, file := range getFiles() {
		switch file.LogFormat {
		case "first_format":
			discoverFirst(c)
		case "second_format":
			discoverSecond(c)
		}
	}

	for e := range c {
		store(e)
	}
}

func getFiles() []model.File {
	panic("the function is not implemented")
}

func discoverFirst(c chan<- model.Entry) {
	panic("the function is not implemented")
}

func discoverSecond(c chan<- model.Entry) {
	panic("the function is not implemented")
}

func store(e model.Entry) {
	panic("the function is not implemented")
}
