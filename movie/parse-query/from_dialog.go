package getscene

import (
	"errors"
	"regexp"
	"strings"

	"github.com/milesq/sw-quotes/config"
	"github.com/milesq/sw-quotes/srt"
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
	MovieID   string
	BegPhrase Phrase
	EndPhrase Phrase
}

// FromDialogQuery .
func FromDialogQuery(rawQuery string, movies []srt.MovieData) (config.ScenePtr, error) {
	var s config.ScenePtr
	quoteWord := `"([^"]+)"(\((\-?[0-9]+)\))?(\[(\d+)\])?`
	quoteRegexp := `^(#(\w+))?` + quoteWord + `\-` + quoteWord + `$`
	quote := regexp.MustCompile(quoteRegexp)

	if !quote.MatchString(rawQuery) {
		return s, ErrNotFound
	}

	query := parseQuery(quote, rawQuery)

	if query.MovieID != "" {
		movies = funk.Filter(movies, func(el srt.MovieData) bool {
			return el.MovieID == query.MovieID
		}).([]srt.MovieData)
	}

	begScene, err := findPhrase(query.BegPhrase, movies)
	if err != nil {
		return s, err
	}

	endScene, err := findPhrase(query.EndPhrase, movies)
	if err != nil {
		return s, err
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

	return s, nil
}

func get_srt_between_scenes(wholeSrts []srt.Subtitle, beg, end srt.Subtitle) []srt.Subtitle {
	return wholeSrts[beg.ID:end.ID]
}
