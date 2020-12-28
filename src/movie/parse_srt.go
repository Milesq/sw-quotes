package movie

import (
	"log"
	"time"

	"github.com/asticode/go-astisub"
	"github.com/thoas/go-funk"
)

// Subtitle contains basic srt data
type Subtitle struct {
	begin time.Duration
	end   time.Duration
	text  string
}

// Parse parses subtitle file
func Parse(name string) []Subtitle {
	srts, err := astisub.OpenFile("./movies/my.srt")

	if err != nil {
		log.Fatal(err)
	}

	return funk.Map(srts.Items, func(s *astisub.Item) Subtitle {
		return Subtitle{s.StartAt, s.EndAt, s.String()}
	}).([]Subtitle)
}
