package srt

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/asticode/go-astisub"
	"github.com/thoas/go-funk"
)

// Subtitle contains basic srt data
type Subtitle struct {
	Begin time.Duration
	End   time.Duration
	Text  string
}

func (s Subtitle) String() string {
	return fmt.Sprintf(
		`%v-%v "%v"`,
		s.Begin.Round(time.Second),
		s.End.Round(time.Second),
		s.Text,
	)
}

// FromFile parses subtitle file
func FromFile(name string) []Subtitle {
	srts, err := astisub.OpenFile(name)

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
