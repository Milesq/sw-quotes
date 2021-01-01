package movie

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/milesq/sw-quotes/src/config"
)

// ScenePtr struct, contains data about search parameter
type ScenePtr struct {
	Name string
	From string
	To   string
}

// NewScenePtr creates Scene Ptr
func NewScenePtr(s string, cfg config.Config) (b ScenePtr) {
	s = strings.ToLower(s)

	// isNamedScene := regexp.MustCompile(`^[0-9a-z/\-\+\\_:]+$`).MatchString(s)

	quoteWord := `"[\w\s]+"(\(\-?[0-9]+\))?(\[\d+\])?`
	quoteRegexp := `^(#\d+)?` + quoteWord + `\-` + quoteWord + `$`
	fmt.Println(quoteRegexp)
	quote := regexp.MustCompile(quoteRegexp)

	matchToQuote := quote.MatchString(s)

	fmt.Printf("%v\n\n", matchToQuote)

	return
}

// func GetScene(scene ScenePtr) {}
