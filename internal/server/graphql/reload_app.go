package graphql

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
)

func (r *mutationResolver) ReloadApp(ctx context.Context, appName string) (*internal.App, error) {
	panic(fmt.Errorf("not implemented"))
}
