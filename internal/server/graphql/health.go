package graphql

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
)

func (r *queryResolver) Health(ctx context.Context) (*internal.Health, error) {
	panic(fmt.Errorf("not implemented"))
}
