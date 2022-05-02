package graphql

import (
	"context"
	"fmt"
)

func (r *mutationResolver) SetAppEnv(ctx context.Context, appName string, newEnv map[string]interface{}) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}
