package fmt

import "fmt"

var cReset = "\x1b[0m"
var cBold = "\x1b[1m"
var cDim = "\x1b[2m"
var cRed = "\x1b[91m"
var cGreen = "\x1b[92m"
var cYellow = "\x1b[93m"
var cBlue = "\x1b[94m"
var cPurple = "\x1b[95m"
var cCyan = "\x1b[96m"

func Println(a ...any) {
	fmt.Println(a...)
}

func Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

func Printfln(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

func Errorf(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

type Logger struct {
}

func (*Logger) V(format string, a ...any) {
	Printfln(cReset+cDim+"[verbose] "+format+cReset, a...)
}

func (*Logger) D(format string, a ...any) {
	Printfln(cReset+"[debug  ] "+format, a...)
}

func (*Logger) I(format string, a ...any) {
	Printfln(cReset+cBlue+"[info   ] "+format+cReset, a...)
}

func (*Logger) W(format string, a ...any) {
	Printfln(cReset+cBold+cYellow+"[warn   ] "+format+cReset, a...)
}

func (*Logger) E(format string, a ...any) {
	Printfln(cReset+cBold+cRed+"[error  ] "+format+cReset, a...)
}
