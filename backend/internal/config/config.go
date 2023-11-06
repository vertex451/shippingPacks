package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string `mapstructure:"SERVER_PORT"`
	LogLevel string `mapstructure:"LOG_LEVEL"`

	PackSize []int `mapstructure:"PACK_SIZE"`
}

func LoadConfig(filePath string) (*Config, error) {
	var err error
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("LOG_LEVEL", "DEBUG")
	viper.SetDefault("PACK_SIZE", []int{250, 5000, 2000, 1000, 500})

	if _, err = os.Stat(filePath); err == nil {
		viper.SetConfigFile(filePath)
		if err = viper.ReadInConfig(); err != nil {
			return nil, err //nolint:wrapcheck
		}
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)

	return &cfg, err //nolint:wrapcheck
}
