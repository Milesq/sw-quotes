package movie

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"

	"github.com/milesq/sw-quotes/src/config"
	parse_query "github.com/milesq/sw-quotes/src/movie/parse-query"
	"github.com/milesq/sw-quotes/src/srt"
	"github.com/thoas/go-funk"
)

// Resolver .
type Resolver struct {
	NamedScenes config.Config
	AllScenes   []srt.MovieData
}

// NewResolver .
func NewResolver(namedScenes config.Config, dir string, movies []string) Resolver {
	moviesWithID := funk.Map(movies, func(s string) [2]string {
		id := md5.Sum([]byte(s))
		return [2]string{s, hex.EncodeToString(id[:])}
	}).([][2]string)

	var allScenes []srt.MovieData

	for _, movie := range moviesWithID {
		var movieData srt.MovieData

		movieData.FileName = movie[0]
		movieData.MovieID = movie[1]

		srtPath := fmt.Sprintf("./%v/%v.srt", dir, movie[0])
		movieData.Srts = srt.FromFile(srtPath)

		allScenes = append(allScenes, movieData)
	}

	return Resolver{
		NamedScenes: namedScenes,
		AllScenes:   allScenes,
	}
}

// Resolve .
func (r *Resolver) Resolve(s string, cfg config.Config) (config.ScenePtr, error) {
	s = strings.ToLower(s)

	isNamedScene := regexp.MustCompile(`^[0-9a-z/\-\+\\_:]+$`).MatchString(s)
	if isNamedScene {
		scene, ok := parse_query.Named(s, cfg)
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
	scene := parse_query.FromDialogQuery(r.AllScenes, query)
	// fmt.Println(r.AllScenes)

	return scene, nil
}
