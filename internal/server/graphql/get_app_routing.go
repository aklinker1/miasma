package graphql

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
)

func (r *queryResolver) GetAppRouting(ctx context.Context, appName string) (*internal.AppRouting, error) {
	panic(fmt.Errorf("not implemented"))
}
