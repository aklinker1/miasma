package graphql

import (
	"context"
	"fmt"
)

func (r *mutationResolver) StopApp(ctx context.Context, appName string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
