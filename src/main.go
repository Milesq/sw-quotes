package main

import (
	"errors"
	"fmt"
	"os"

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
	// movie.NewScenePtr(`"You turned her against me"-"I will do what I must"`, globalCfg)
	// movie.NewScenePtr(`#4"You turned her against me"(-2)[2]-"I will do what I must"(3)[3]`, globalCfg)
	s, err := movie.NewScenePtr(`luke:hallway`, globalCfg)
	fmt.Println(s)
	fmt.Println(errors.Is(err, os.ErrNotExist))
}
