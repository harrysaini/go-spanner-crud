package models

type gcp struct {
	Project  string `mapstructure:"project"`
	Instance string `mapstructure:"instance"`
	Database string `mapstructure:"database"`
}

type server struct {
	Port int `mapstructure:"port"`
}

type app struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type redis struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	TTL  int64  `mapstructure:"ttl"`
}

// Configuration - model for app conf
type Configuration struct {
	Gcp    gcp    `mapstructure:"gcp"`
	App    app    `mapstructure:"app"`
	Server server `mapstructure:"server"`
	Redis  redis  `mapstructure:"redis"`
}
