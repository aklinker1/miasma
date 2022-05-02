package server

type Logger interface {
	D(format string, args ...any)
	V(format string, args ...any)
	I(format string, args ...any)
	W(format string, args ...any)
	E(format string, args ...any)
}
