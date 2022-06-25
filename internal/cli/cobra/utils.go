package cobra

import "github.com/fatih/color"

var yellow = color.New(color.FgHiYellow).SprintFunc()
var green = color.New(color.FgHiGreen).SprintFunc()
var red = color.New(color.FgHiRed).SprintFunc()
var purple = color.New(color.FgHiMagenta).SprintFunc()
var bold = color.New(color.Bold).SprintFunc()
var dim = color.New(color.Faint).SprintFunc()

var title = color.New(color.FgHiMagenta, color.Bold)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
