package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

// ReadFiles .
func ReadFiles(dir string) (ret []string) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		parts := strings.Split(file.Name(), ".")
		endCursor := len(parts) - 1
		name, ext := parts[:endCursor], parts[endCursor]

		if ext == "srt" {
			ret = append(ret, strings.Join(name, "."))
		}
	}

	return
}
