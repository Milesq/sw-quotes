package getscene

import (
	"errors"
	"fmt"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/milesq/sw-quotes/src/srt"
)

const phrasePassThreshold = .85

func findPhrase(phrase string, movies []srt.MovieData) (srt.MovieData, error) {
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

	switch len(foundedScene) {
	case 1:
		return foundedScene[0], nil
	case 0:
		return srt.MovieData{}, errors.New("we couldnt find the following scene: " + phrase)
	default:
		err := "we found multiple matching phrases. You have to precise which one you want to use\n"
		for i, scene := range foundedScene {
			err += fmt.Sprint(
				"Scene ID: ",
				i,
				"\n\tMovie Name: ",
				scene.FileName,
				"\n\tMovie ID: ",
				scene.MovieID,
				"\n\tPhrase: ",
				scene.Srts[0],
				"\n\n",
			)
		}

		return srt.MovieData{}, errors.New(err)
	}
}
