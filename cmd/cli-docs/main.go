package main

import (
	"fmt"
	"strings"

	"github.com/aklinker1/miasma/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	printCommand(cli.RootCmd, "##")
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
