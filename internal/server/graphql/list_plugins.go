package graphql

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
)

func (r *queryResolver) ListPlugins(ctx context.Context) ([]internal.Plugin, error) {
	panic(fmt.Errorf("not implemented"))
}
