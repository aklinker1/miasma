package log

import (
	"fmt"
	"os"
)

var RESET = "\x1b[0m"
var DIM = "\x1b[2m"
var RED = "\x1b[91m"
var YELLOW = "\x1b[93m"
var BLUE = "\x1b[94m"
var GREEN = "\x1b[92m"

func printColored(color, format, messageFormat string, args ...interface{}) {
	message := fmt.Sprintf(messageFormat, args...)
	fmt.Printf("%s"+format+"%s\n", color, message, RESET)
}

func V(format string, args ...interface{}) {
	printColored(DIM, "[ verbose ] %s", format, args...)
}

func D(format string, args ...interface{}) {
	printColored(GREEN, "[ debug   ] %s", format, args...)
}

func I(format string, args ...interface{}) {
	printColored(BLUE, "[ info    ] %s", format, args...)
}

func W(format string, args ...interface{}) {
	printColored(YELLOW, "[ warning ] %s", format, args...)
}

func E(format string, args ...interface{}) {
	printColored(RED, "[ error   ] %s", format, args...)
}

func Fatal(format string, args ...interface{}) {
	printColored(RED, "[ fatal   ] %s", format, args...)
	os.Exit(1)
}
