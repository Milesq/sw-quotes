package movie

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"

	"github.com/milesq/sw-quotes/config"
	parse_query "github.com/milesq/sw-quotes/movie/parse-query"
	"github.com/milesq/sw-quotes/srt"
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
	var (
		scene config.ScenePtr
		err   error
	)

	isNamedScene := regexp.MustCompile(`^[0-9a-z/\-\+\\_:]+$`).MatchString(s)

	if isNamedScene {
		scene, err = parse_query.Named(s, cfg)
	} else {
		scene, err = parse_query.FromDialogQuery(s, r.AllScenes)
	}

	if err != nil {
		return config.ScenePtr{}, err
	}

	return scene, nil
}
