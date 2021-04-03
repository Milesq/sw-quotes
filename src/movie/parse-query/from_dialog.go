package getscene

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/milesq/sw-quotes/src/config"
	"github.com/milesq/sw-quotes/src/srt"
	"github.com/thoas/go-funk"
)

// Phrase .
type Phrase struct {
	Str    string
	Offset int
	I      int
}

// Query .
type Query struct {
	MovieID   int
	BegPhrase Phrase
	EndPhrase Phrase
}

// FromDialogQuery .
func FromDialogQuery(rawQuery string, scenes []srt.MovieData) (s config.ScenePtr, nil error) {
	quoteWord := `"([\w\s]+)"(\((\-?[0-9]+)\))?(\[(\d+)\])?`
	quoteRegexp := `^(#(\d+))?` + quoteWord + `\-` + quoteWord + `$`
	quote := regexp.MustCompile(quoteRegexp)

	if !quote.MatchString(rawQuery) {
		return s, ErrNotFound
	}

	query := parseQuery(quote, rawQuery)
	scene, err := findPhrase(query.BegPhrase.Str, scenes)

	if err != nil {
		fmt.Println(err)
	}

	s.Filename = scene.FileName

	s.Srt = strings.Join(funk.Map(scene.Srts, func(el srt.Subtitle) string {
		return el.Text
	}).([]string), " ")

	s.Timestamp = [][2]string{
		{
			scene.Srts[0].Begin.String(),
			scene.Srts[len(scene.Srts)-1].End.String(),
		},
	}

	return
}
