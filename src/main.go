package main

import (
	"fmt"

	"github.com/milesq/sw-quotes/src/movie"
)

func main() {
	fmt.Println(movie.Parse("movies/my"))
}
