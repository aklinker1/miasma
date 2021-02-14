package main

import (
	"fmt"
	"strings"

	"github.com/aklinker1/miasma/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	fmt.Println(strings.Join([]string{
		"---",
		"id: cli-usage",
		"title: CLI Usage",
		"description: CLI commands and documentation",
		"slug: /cli",
		"---",
		"",
		"The Misama CLI should be installed on any computer that you would deploy/manage application from. It's very similar to Heroku's CLI, but also takes inspiration from the Docker CLI.",
		"",
		"Checkout the [Get Started](/docs#install-cli) page to learn how to install the CLI",
		"",
	}, "\n"))
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
