package main

import (
	"github.com/milesq/sw-quotes/config"
	"github.com/milesq/sw-quotes/movie"
	"github.com/milesq/sw-quotes/utils"
)

const movieDir = "movies"

var resolver movie.Resolver

func init() {
	globalCfg, err := config.ParseFromFile("./movies.config.yml")

	if err != nil {
		panic("cannot parse from file")
	}

	resolver = movie.NewResolver(globalCfg, movieDir, utils.ReadFiles(movieDir))
}

// Resolve .
func Resolve(query string) (config.ScenePtr, error) {
	return resolver.Resolve(query)
}
