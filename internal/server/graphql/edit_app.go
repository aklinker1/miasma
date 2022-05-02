package graphql

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
)

func (r *mutationResolver) EditApp(ctx context.Context, appName string, app internal.EditAppInput) (*internal.App, error) {
	panic(fmt.Errorf("not implemented"))
}
