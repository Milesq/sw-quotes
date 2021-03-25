package getscene

import (
	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/milesq/sw-quotes/src/srt"
)

const phrasePassThreshold = .85

func findPhrase(phrase string, movies []srt.MovieData) error {
	var foundedScene []srt.MovieData

	for _, movie := range movies {
		for _, subtitle := range movie.Srts {
			sm := strutil.Similarity(subtitle.Text, phrase, metrics.NewJaro())

			if sm >= phrasePassThreshold {
				scene := movie
				scene.Srts = []srt.Subtitle{subtitle}
				foundedScene = append(foundedScene, scene)
			}
		}
	}

	return nil
}
