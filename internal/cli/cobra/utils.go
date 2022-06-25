package cobra

import (
	"fmt"

	"github.com/fatih/color"
)

var yellow = color.New(color.FgHiYellow).SprintFunc()
var green = color.New(color.FgHiGreen).SprintFunc()
var red = color.New(color.FgHiRed).SprintFunc()
var purple = color.New(color.FgHiMagenta).SprintFunc()
var bold = color.New(color.Bold).SprintFunc()
var dim = color.New(color.Faint).SprintFunc()

var title = color.New(color.FgHiMagenta, color.Bold)
var fatal = color.New(color.FgHiRed, color.Bold)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func done(format string, args ...any) {
	coloredFormat := fmt.Sprintf(green("✔ %s\n\n"), format)
	fmt.Printf(coloredFormat, args...)
}
