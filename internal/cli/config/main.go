package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	miasmaAPI "github.com/aklinker1/miasma/package/client"
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

func Client() *miasmaAPI.Miasma {
	host := viper.GetString("host")
	if host == "" {
		fmt.Println("Miasma CLI is not connected to a server yet. Run 'miasma connect <ip:port>'")
		os.Exit(1)
	}
	client := miasmaAPI.NewClientWith(host)
	return client
}
