package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/milesq/sw-quotes/src/config"
	"github.com/milesq/sw-quotes/src/movie"
)

const movieDir = "movies"

func main() {
	globalCfg, err := config.ParseFromFile("./movies.config.yml")

	if err != nil {
		panic("cannot parse from file")
	}

	resolver := movie.NewResolver(globalCfg, movieDir, readFiles(movieDir))

	resolver.Resolve(`luked:hallway`, globalCfg)
	// resolver.Resolve(`"You turned her against me"-"I will do what I must"`, globalCfg)
	// resolver.Resolve(`#1"You turned her against me"(-2)[3]-"I will do what I must"(4)[5]`, globalCfg)
}

func readFiles(dir string) (ret []string) {
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
