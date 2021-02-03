package models

type GCP struct {
	Project  string `mapstructure:"project"`
	Instance string `mapstructure:"instance"`
	Database string `mapstructure:"database"`
}

type Server struct {
	Port int `mapstructure:"port"`
}

type App struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type Configuration struct {
	Gcp    GCP    `mapstructure:"gcp"`
	App    App    `mapstructure:"app"`
	Server Server `mapstructure:"server"`
}
