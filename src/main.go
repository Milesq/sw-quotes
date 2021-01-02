package main

import (
	"github.com/milesq/sw-quotes/src/config"
)

const movieDir = "movies"

func main() {
	globalCfg, err := config.ParseFromFile("./movies.config.yml")

	if err != nil {
		panic("cannot parse from file")
	}

	// fmt.Println(config.ParseFromFile("./movies.config.yml"))
	// movie.NewScenePtr(`"You turned her against me"-"I will do what I must"`, globalCfg)
	// movie.NewScenePtr(`#4"You turned her against me"(-2)[2]-"I will do what I must"(3)[3]`, globalCfg)
	// movie.NewScenePtr(`luked:hallway`, globalCfg)
}
