package config

import (
	"sync"

	"github.com/spf13/viper"
)

// AppConfig represents the application config that are defined in env files.
type AppConfig struct {
	Port string `mapstructure:"port"`
}

// Config is ...
var Config AppConfig

// LoadConfig loads config file and marshals in Config var
func LoadConfig(path string) {
	var configOnce sync.Once

	configOnce.Do(func() {
		env := GetEnvironment()

		viper.SetConfigName(env + ".config")
		viper.AddConfigPath(path)
		viper.SetConfigType("yaml")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		err := viper.Unmarshal(&Config)

		if err != nil {
			panic(err)
		}
	})
}
