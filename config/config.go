package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	var err error
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")   // optionally look for config in the working directory
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func GetString(s string) string {
	return viper.GetString(s)
}

func GetInt(s string) int {
	return viper.GetInt(s)
}

func GetBool(s string) bool {
	return viper.GetBool(s)
}

func GetFloat64(s string) float64 {
	return viper.GetFloat64(s)
}
