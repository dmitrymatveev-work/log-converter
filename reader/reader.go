package reader

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Reader handles file reading with offset memory
type Reader struct {
	files map[string]int64
}

// New creates new instance of Reader
func New() *Reader {
	return &Reader{
		files: make(map[string]int64),
	}
}

func (r *Reader) Read(path string) []string {
	var off int64

	if o, ok := r.files[path]; ok {
		off = o
	} else {
		r.files[path] = off
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	file.Seek(off, 0)

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	r.files[path] = off + int64(len(bytes))

	str := string(bytes)

	return strings.Split(str, "\n")
}
