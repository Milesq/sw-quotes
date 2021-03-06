package getscene

import (
	"errors"
	"fmt"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/milesq/sw-quotes/srt"
)

const phrasePassThreshold = .85

func findPhrase(phrase Phrase, movies []srt.MovieData) (srt.MovieData, error) {
	var foundedScene []srt.MovieData

	for _, movie := range movies {
		for _, subtitle := range movie.Srts {
			sm := strutil.Similarity(subtitle.Text, phrase.Str, metrics.NewJaro())

			if sm >= phrasePassThreshold {
				scene := movie
				scene.Srts = []srt.Subtitle{subtitle}
				foundedScene = append(foundedScene, scene)
			}
		}
	}

	if phrase.I != -1 {
		if len(foundedScene) <= phrase.I {
			err := fmt.Sprintln("There is only ", len(foundedScene), " scenes")
			for _, scene := range foundedScene {
				err += fmt.Sprintf("`%v`\n", scene.Srts[0].Text)
				err += scene.Srts[0].Begin.String() + "\n\n"
			}

			return srt.MovieData{}, errors.New(err)
		}
		foundedScene = []srt.MovieData{foundedScene[phrase.I]}
	}

	switch len(foundedScene) {
	case 1:
		return foundedScene[0], nil
	case 0:
		return srt.MovieData{}, errors.New("we couldnt find the following scene: " + phrase.Str)
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
