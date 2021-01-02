package movie

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/thoas/go-funk"

	"github.com/milesq/sw-quotes/src/config"
)

// NewScenePtr creates Scene Ptr
func NewScenePtr(s string, cfg config.Config) (config.ScenePtr, error) {
	s = strings.ToLower(s)

	isNamedScene := regexp.MustCompile(`^[0-9a-z/\-\+\\_:]+$`).MatchString(s)
	if isNamedScene {
		getNamedScene(s, cfg)
		scene, ok := getNamedScene(s, cfg)
		if !ok {
			return config.ScenePtr{}, os.ErrNotExist
		}

		return scene, nil
	}

	quoteWord := `"[\w\s]+"(\(\-?[0-9]+\))?(\[\d+\])?`
	quoteRegexp := `^(#\d+)?` + quoteWord + `\-` + quoteWord + `$`
	quote := regexp.MustCompile(quoteRegexp)

	matchToQuote := quote.MatchString(s)

	fmt.Printf("%v\n\n", matchToQuote)

	return config.ScenePtr{}, nil
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
