package graphql

import (
	"context"

	"github.com/aklinker1/miasma/internal/server"
)

// Returns fallback and the err if it exists, otherwise it returns the value and nil
func safeReturn[T any](value T, fallback T, err error) (T, error) {
	if err != nil {
		return fallback, err
	} else {
		return value, err
	}
}

// Runs a function block inside a transaction. If the function returns no error, the result is
// committed, otherwise the transaction is rolled back
func inTx[T interface{}](ctx context.Context, beginTx func(context.Context) (server.Tx, error), defaultValue T, fn func(server.Tx) (T, error)) (T, error) {
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
