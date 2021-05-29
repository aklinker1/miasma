package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	miasmaAPI "github.com/aklinker1/miasma/package/client"
	"github.com/aklinker1/miasma/package/client/operations"
)

var connectCmd = &cobra.Command{
	Use:     "connect",
	Short:   "Connect to a Miasma Server",
	Args:    cobra.ExactArgs(1),
	Example: "miasma connect localhost:3000",
	Run: func(cmd *cobra.Command, args []string) {
		connectToServer(args[0])
	},
}

func init() {
	RootCmd.AddCommand(connectCmd)
}

func connectToServer(host string) {
	client := miasmaAPI.NewClientWith(host)
	viper.Set("host", host)

	_, err := client.Operations.HealthCheck(operations.NewHealthCheckParams())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Connected to miasma!")
	err = viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s added to %s\n", host, viper.ConfigFileUsed())
}
