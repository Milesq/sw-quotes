package movie

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/milesq/sw-quotes/src/config"
	get_scene "github.com/milesq/sw-quotes/src/movie/get-scene"
)

// NewScenePtr creates Scene Ptr
func NewScenePtr(s string, cfg config.Config) (config.ScenePtr, error) {
	s = strings.ToLower(s)

	isNamedScene := regexp.MustCompile(`^[0-9a-z/\-\+\\_:]+$`).MatchString(s)
	if isNamedScene {
		scene, ok := get_scene.Named(s, cfg)
		if !ok {
			return config.ScenePtr{}, ErrNotFound
		}

		return scene, nil
	}

	quoteWord := `"([\w\s]+)"(\((\-?[0-9]+)\))?(\[(\d+)\])?`
	quoteRegexp := `^(#(\d+))?` + quoteWord + `\-` + quoteWord + `$`
	quote := regexp.MustCompile(quoteRegexp)

	if !quote.MatchString(s) {
		return config.ScenePtr{}, ErrQueryDoesntMatch
	}

	query := parseQuery(quote, s)
	scene := get_scene.FromQuery(query, cfg)

	return config.ScenePtr{}, nil
}

func parseQuery(r *regexp.Regexp, s string) get_scene.Query {
	atoi := func(str string) int {
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}

		return num
	}

	parts := r.FindStringSubmatch(s)

	return query{
		atoi(parts[2]),
		phrase{parts[3], atoi(parts[5]), atoi(parts[7])},
		phrase{parts[8], atoi(parts[10]), atoi(parts[12])},
	}
}

// func GetScene(scene ScenePtr) {}
