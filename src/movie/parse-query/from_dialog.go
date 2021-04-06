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
func FromDialogQuery(rawQuery string, movies []srt.MovieData) (s config.ScenePtr, nil error) {
	quoteWord := `"([\w\s]+)"(\((\-?[0-9]+)\))?(\[(\d+)\])?`
	quoteRegexp := `^(#(\d+))?` + quoteWord + `\-` + quoteWord + `$`
	quote := regexp.MustCompile(quoteRegexp)

	if !quote.MatchString(rawQuery) {
		return s, ErrNotFound
	}

	query := parseQuery(quote, rawQuery)
	begScene, err := findPhrase(query.BegPhrase.Str, movies)
	if err != nil {
		fmt.Println(err)
	}

	endScene, err := findPhrase(query.EndPhrase.Str, movies)
	if err != nil {
		fmt.Println(err)
	}

	if begScene.MovieID != endScene.MovieID {
		return s, errors.New("scenes doesnt come from the same movie")
	}

	s.Filename = begScene.FileName

	movieScenes := funk.Find(movies, func(movie srt.MovieData) bool {
		return movie.MovieID == begScene.MovieID
	}).(srt.MovieData).Srts
	wholeSceneSrts := get_srt_between_scenes(movieScenes, begScene.Srts[0], endScene.Srts[0])

	s.Srt = strings.Join(funk.Map(wholeSceneSrts, func(el srt.Subtitle) string {
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

func get_srt_between_scenes(wholeSrts []srt.Subtitle, beg, end srt.Subtitle) []srt.Subtitle {
	return wholeSrts[beg.ID:end.ID]
}
