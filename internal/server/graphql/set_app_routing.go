package graphql

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
)

func (r *mutationResolver) SetAppRouting(ctx context.Context, appName string, routing *internal.AppRoutingInput) (*internal.AppRouting, error) {
	panic(fmt.Errorf("not implemented"))
}
