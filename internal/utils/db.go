package utils

import (
	"context"

	"github.com/aklinker1/miasma/internal/server"
)

// Runs a function block inside a transaction. If the function returns no error, the result is
// committed, otherwise the transaction is rolled back
func InTx[T interface{}](ctx context.Context, beginTx func(context.Context) (server.Tx, error), defaultValue T, fn func(server.Tx) (T, error)) (T, error) {
	tx, err := beginTx(ctx)
	if err != nil {
		return defaultValue, err
	}
	defer tx.Rollback()

	v, err := fn(tx)
	if err != nil {
		return defaultValue, err
	}
	tx.Commit()
	return v, nil
}
