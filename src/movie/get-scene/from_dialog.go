package getscene

import "github.com/milesq/sw-quotes/src/config"

// Phrase .
type Phrase struct {
	str    string
	offset int
	i      int
}

// Query .
type Query struct {
	MovieID   int
	BegPhrase Phrase
	EndPhrase Phrase
}

// FromQuery .
func FromQuery(query Query, cfg config.Config) config.ScenePtr {
	return config.ScenePtr{}
}
