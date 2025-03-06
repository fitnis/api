package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPass     string `mapstructure:"DB_PASS"`
	DBName     string `mapstructure:"DB_NAME"`
	ServerPort string `mapstructure:"SERVER_PORT"`
}

func LoadConfig() (Config, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // or "toml", "json" ...
	viper.AddConfigPath(".")      // look for config in the current directory
	// viper.AddConfigPath("/etc/fitnis/") // or additional paths

	// Override config with environment variables if present
	viper.AutomaticEnv()

	// Try reading config file
	// If no file is found, we'll just use environment variables
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("No config file found or error reading config file: %v\nUsing environment variables only.", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
