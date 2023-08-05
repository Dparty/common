package notifications

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestSend(t *testing.T) {
	var err error
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")   // optionally look for config in the working directory
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("configuration error: %w", err))
	}
}
