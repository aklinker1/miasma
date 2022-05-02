package server

type TXMode int

const (
	ReadOnly  TXMode = 0
	ReadWrite TXMode = 1
)

type TX interface {
	Commit()
	Rollback()
}

type DB interface {
	Open() error
	TX(mode TXMode) (TX, error)
}
