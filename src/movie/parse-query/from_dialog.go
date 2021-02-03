package getscene

import (
	"fmt"

	"github.com/milesq/sw-quotes/src/config"
	"github.com/milesq/sw-quotes/src/srt"
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
func FromDialogQuery(scenes []srt.MovieData, query Query) config.ScenePtr {
	fmt.Println(query)
	return config.ScenePtr{}
}
