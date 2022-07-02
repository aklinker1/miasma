package cobra

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/aklinker1/miasma/internal/cli"
	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var metadata cli.Metadata
var api cli.APIService

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "miasma",
	Short: "Manage and deploy dockerized applications to a docker swarm",
	Run: func(cmd *cobra.Command, args []string) {
		if flags.GetVersionFlag(cmd) {
			fmt.Printf("Miasma CLI version %s, build %s\n", metadata.Version, metadata.BuildHash)

			// Don't execute any other command
			os.Exit(0)
		}
	},
}

func Execute(
	injectedMetadata cli.Metadata,
	injectedAPI cli.APIService,
) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf(red(bold("Fatal error: %v\n\n")), err)
			fmt.Println(dim(string(debug.Stack())))
			os.Exit(1)
		}
	}()

	metadata = injectedMetadata
	api = injectedAPI

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.miasma.yaml)")
	flags.UseVersionFlag(RootCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".miasma" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".miasma")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match
	config.ReadConfig()
	api.SetBaseURL(viper.GetString("host"))
	api.SetAccessToken(viper.GetString("accessToken"))
}
