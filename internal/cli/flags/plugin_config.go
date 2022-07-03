package flags

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UsePluginConfigFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("plugin-config", "c", "", "JSON string representing the plugin's config. Example: --plugin-config '{ \"key\": \"value\" }'")
}

func GetPluginConfigFlag(cmd *cobra.Command) map[string]any {
	v, err := cmd.Flags().GetString("plugin-config")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config := map[string]any{}
	if v != "" {
		json.Unmarshal([]byte(v), &config)
	}
	return config
}
