package config

// ScenePtr struct, contains data about search parameter
type ScenePtr struct {
	Name      string
	Filename  string
	Srt       string
	Timestamp [][2]string
}

// Config contains scenes from config file
type Config []ScenePtr
