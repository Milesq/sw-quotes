package getscene

import (
	"errors"
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
	begScene, err := findPhrase(query.BegPhrase.Str, scenes)
	if err != nil {
		fmt.Println(err)
	}

	endScene, err := findPhrase(query.EndPhrase.Str, scenes)
	if err != nil {
		fmt.Println(err)
	}

	if begScene.FileName != endScene.FileName {
		return s, errors.New("scenes doesnt come from the same movie")
	}

	s.Filename = begScene.FileName

	s.Srt = strings.Join(funk.Map(begScene.Srts, func(el srt.Subtitle) string {
		return el.Text
	}).([]string), " ")

	s.Timestamp = [][2]string{
		{
			begScene.Srts[0].Begin.String(),
			endScene.Srts[len(endScene.Srts)-1].End.String(),
		},
	}

	return
}
