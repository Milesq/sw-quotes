package eventHandlers

import (
	"strings"

	"github.com/milesq/sw-quotes/config"
	"github.com/milesq/sw-quotes/movie"
	"github.com/milesq/sw-quotes/utils"
	"github.com/thoas/go-funk"
)

const movieDir = "movies"

var resolver movie.Resolver
var globalConfig config.Config
var predefinedSceneInfo string

func getPredefinedScenes(globalConfig config.Config) string {
	result := "There are defined following scenes:\n"

	predefinedSceneNames := funk.Map(
		globalConfig,
		func(el config.ScenePtr) string {
			return el.Name
		},
	).([]string)

	sceneSeparator := "\n\t- "

	result += sceneSeparator
	result += strings.Join(predefinedSceneNames, sceneSeparator)

	return result
}

func init() {
	globalConfig, _ = config.ParseFromFile("./movies.config.yml")

	resolver = movie.NewResolver(globalConfig, movieDir, utils.ReadFiles(movieDir))
	predefinedSceneInfo = getPredefinedScenes(globalConfig)
}

func resolveQuery(query string) (config.ScenePtr, error) {
	return resolver.Resolve(query)
}
