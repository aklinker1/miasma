package flags

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func UseSetTraefikFlags(cmd *cobra.Command) {
	cmd.Flags().String("host", "", "The hostname of the app. EX: test.home.io")
	cmd.Flags().String("path", "", "The path at the host the app will live at. EX: /api")
	cmd.Flags().String("rule", "", "Custom traefik routing rule")
}

type SetTraefik struct {
	Host *string
	Path *string
	Rule *string
}

func GetSetTraefikFlags(cmd *cobra.Command) *SetTraefik {
	host, err := cmd.Flags().GetString("host")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	hostTrimmed := strings.TrimSpace(host)
	var hostPtr *string
	if hostTrimmed != "" {
		hostPtr = &hostTrimmed
	}

	path, err := cmd.Flags().GetString("path")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pathTrimmed := strings.TrimSpace(path)
	var pathPtr *string
	if pathTrimmed != "" {
		pathPtr = &pathTrimmed
	}

	rule, err := cmd.Flags().GetString("rule")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ruleTrimmed := strings.TrimSpace(rule)
	var rulePtr *string
	if ruleTrimmed != "" {
		rulePtr = &ruleTrimmed
	}

	return &SetTraefik{
		Host: hostPtr,
		Path: pathPtr,
		Rule: rulePtr,
	}
}
