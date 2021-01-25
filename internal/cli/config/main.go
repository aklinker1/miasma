package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func ReadConfig() {
	err := viper.ReadInConfig()
	if err == nil {
		return
	}
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		err := viper.SafeWriteConfig()
		if err != nil {
			fmt.Println("Failed to initialize config file")
			os.Exit(1)
		}
		return
	}
}
