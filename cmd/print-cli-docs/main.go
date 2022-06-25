package main

import (
	"fmt"
	"strings"

	cobra2 "github.com/aklinker1/miasma/internal/cli/cobra"
	"github.com/spf13/cobra"
)

func main() {
	printCommand(cobra2.RootCmd, "##")
}

func printCommand(cmd *cobra.Command, headerPrefix string) {
	if len(headerPrefix) > 6 {
		panic("docs cannot be generated for sub-commands cannot deeper than level 6")
	}
	fmt.Println(strings.Join([]string{
		fmt.Sprintf("%s `%s`", headerPrefix, cmd.Name()),
		"",
		"```",
	}, "\n"))
	cmd.Help()
	fmt.Println(strings.Join([]string{
		"```",
		"",
	}, "\n"))

	for _, childCmd := range cmd.Commands() {
		printCommand(childCmd, headerPrefix+"#")
	}
}
