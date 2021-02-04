package libs

import (
	"go-spanner-crud/src/models"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Conf - config of app
var Conf *models.Configuration

func init() {
	var err error

	// Loads relative to main.go
	Conf, err = loadConfig("./config")

	if err != nil {
		log.Println("Failed to load conf")
		log.Fatalln(err)
	}
}

func loadConfig(path string) (*models.Configuration, error) {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.ReadInConfig()

	env := "dev"
	if envVar := os.Getenv("ENV"); envVar != "" {
		env = strings.ToLower(envVar)
	}

	viper.SetConfigName(env)
	viper.MergeInConfig()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var conf models.Configuration
	err := viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
