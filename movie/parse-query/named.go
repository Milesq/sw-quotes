package getscene

import (
	"github.com/thoas/go-funk"

	"github.com/milesq/sw-quotes/config"
)

// Named .
func Named(name string, cfg config.Config) (config.ScenePtr, error) {
	found := funk.Find(cfg, func(sc config.ScenePtr) bool {
		return sc.Name == name
	})

	if found == nil {
		return config.ScenePtr{}, ErrNotFound
	}

	return found.(config.ScenePtr), nil
}
