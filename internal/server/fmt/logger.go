package fmt

import "fmt"

func Println(a ...any) {
	fmt.Println(a...)
}

func Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

func Printfln(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

type Logger struct {
}

func (*Logger) V(format string, a ...any) {
	Printfln("[verbose] "+format, a...)
}

func (*Logger) D(format string, a ...any) {
	Printfln("[debug  ] "+format, a...)
}

func (*Logger) I(format string, a ...any) {
	Printfln("[info   ] "+format, a...)
}

func (*Logger) W(format string, a ...any) {
	Printfln("[warn   ] "+format, a...)
}

func (*Logger) E(format string, a ...any) {
	Printfln("[error  ] "+format, a...)
}
