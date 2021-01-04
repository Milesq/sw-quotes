package main

import (
	"github.com/milesq/sw-quotes/src/config"
	"github.com/milesq/sw-quotes/src/movie"
)

const movieDir = "movies"

func main() {
	globalCfg, err := config.ParseFromFile("./movies.config.yml")

	if err != nil {
		panic("cannot parse from file")
	}

	// fmt.Println(config.ParseFromFile("./movies.config.yml"))
	movie.NewScenePtr(`"You turned her against me"-"I will do what I must"`, globalCfg)
	// movie.NewScenePtr(`#1"You turned her against me"(-2)[3]-"I will do what I must"(4)[5]`, globalCfg)
	// movie.NewScenePtr(`luked:hallway`, globalCfg)
}
