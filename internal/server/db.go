package server

import (
	"context"
	"database/sql"
)

type Tx = *sql.Tx

type DB interface {
	Open() error
	ReadonlyTx(ctx context.Context) (Tx, error)
	ReadWriteTx(ctx context.Context) (Tx, error)
}
