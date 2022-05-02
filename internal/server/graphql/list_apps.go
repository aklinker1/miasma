package graphql

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
)

func (r *queryResolver) ListApps(ctx context.Context, page *int32, size *int32, showHidden *bool) ([]internal.App, error) {
	panic(fmt.Errorf("not implemented"))
}
