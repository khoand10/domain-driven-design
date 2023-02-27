package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	APIVersion string `mapstructure:"API_VERSION"`
	SqlitePath string `mapstructure:"DB_SQLITE_PATH"`
}

var config *Config

func LoadConfig(path string) *Config {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	err = viper.Unmarshal(&config)
	return config
}

func GetConfig() *Config {
	return config
}
