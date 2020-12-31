package movie

import (
	"log"
	"regexp"
	"strings"
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

// ParseFromFile parses subtitle file
func ParseFromFile(name string) []Subtitle {
	srts, err := astisub.OpenFile("./movies/my.srt")

	if err != nil {
		log.Fatal(err)
	}

	return funk.Map(srts.Items, func(s *astisub.Item) Subtitle {
		return Subtitle{s.StartAt, s.EndAt, normalize(s.String())}
	}).([]Subtitle)
}

func normalize(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

	return strings.ToLower(re.ReplaceAllString(s, ""))
}
