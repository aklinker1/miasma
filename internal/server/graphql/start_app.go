package graphql

import (
	"context"
	"fmt"
)

func (r *mutationResolver) StartApp(ctx context.Context, appName string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
