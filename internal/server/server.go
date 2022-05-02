package server

type Server interface {
	ServeGraphql() error
}
