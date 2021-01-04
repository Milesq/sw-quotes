package movie

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"

	"github.com/milesq/sw-quotes/src/config"
)

type phrase struct {
	str    string
	offset int
	i      int
}

type query struct {
	movieID   int
	begPhrase phrase
	endPhrase phrase
}

// NewScenePtr creates Scene Ptr
func NewScenePtr(s string, cfg config.Config) (config.ScenePtr, error) {
	s = strings.ToLower(s)

	isNamedScene := regexp.MustCompile(`^[0-9a-z/\-\+\\_:]+$`).MatchString(s)
	if isNamedScene {
		getNamedScene(s, cfg)
		scene, ok := getNamedScene(s, cfg)
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
	fmt.Printf("%v", query)

	return config.ScenePtr{}, nil
}

func parseQuery(r *regexp.Regexp, s string) query {
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

func getNamedScene(name string, cfg config.Config) (s config.ScenePtr, false bool) {
	found := funk.Find(cfg, func(sc config.ScenePtr) bool {
		return sc.Name == name
	})

	if found == nil {
		return
	}

	return found.(config.ScenePtr), true
}

// func GetScene(scene ScenePtr) {}
