package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func init() {
	var err error
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")   // optionally look for config in the working directory
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("databases fatal error config file: %w", err))
	}
}

func GetString(s string) string {
  if value := viper.GetString(s); value != "" {
    return value
  }
  if value := os.Getenv(s); value != "" {
    return value
  }
  return ""
}
