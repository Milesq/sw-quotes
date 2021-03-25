package getscene

import (
	"fmt"
	"regexp"

	"github.com/milesq/sw-quotes/src/config"
	"github.com/milesq/sw-quotes/src/srt"
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
	err := findPhrase(query.BegPhrase.Str, scenes)

	if err != nil {
		fmt.Println(err)
	}

	return
}
