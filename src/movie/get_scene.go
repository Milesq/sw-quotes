package movie

import (
	"regexp"
	"strconv"

	get_scene "github.com/milesq/sw-quotes/src/movie/get-scene"
)

func parseQuery(r *regexp.Regexp, s string) get_scene.Query {
	atoi := func(str string) int {
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}

		return num
	}

	parts := r.FindStringSubmatch(s)

	return get_scene.Query{
		atoi(parts[2]),
		get_scene.Phrase{parts[3], atoi(parts[5]), atoi(parts[7])},
		get_scene.Phrase{parts[8], atoi(parts[10]), atoi(parts[12])},
	}
}

// func GetScene(scene ScenePtr) {}
