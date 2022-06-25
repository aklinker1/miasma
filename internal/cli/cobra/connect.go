package cobra

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func connectToServer(baseURL string) {
	ctx := context.Background()
	api.SetBaseURL(baseURL)
	viper.Set("host", baseURL)

	health, err := api.Health(ctx, `{
		version
		dockerVersion
		cluster {
			id
			joinCommand
			createdAt
			updatedAt
		}
	}`)
	checkErr(err)

	title.Println("\nConnected to miasma!")
	fmt.Println(dim("  Miasma Server Version ", health.Version))
	fmt.Println(dim("  Docker Version ", health.DockerVersion))
	fmt.Println(dim("  Cluster Enabled? ", health.Cluster != nil))

	err = viper.WriteConfig()
	checkErr(err)
	fmt.Printf("%s added to %s\n\n", baseURL, viper.ConfigFileUsed())
}
