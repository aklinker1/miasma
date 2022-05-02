package graphql

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
)

func (r *mutationResolver) CreateApp(ctx context.Context, app internal.CreateAppInput) (*internal.App, error) {
	panic(fmt.Errorf("not implemented"))
}
