package graphql

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
)

func (r *mutationResolver) RemoveAppRouting(ctx context.Context, appName string) (*internal.AppRouting, error) {
	panic(fmt.Errorf("not implemented"))
}
