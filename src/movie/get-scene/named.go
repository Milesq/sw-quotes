package getscene

import (
	"github.com/thoas/go-funk"

	"github.com/milesq/sw-quotes/src/config"
)

// Named .
func Named(name string, cfg config.Config) (s config.ScenePtr, false bool) {
	found := funk.Find(cfg, func(sc config.ScenePtr) bool {
		return sc.Name == name
	})

	if found == nil {
		return
	}

	return found.(config.ScenePtr), true
}
