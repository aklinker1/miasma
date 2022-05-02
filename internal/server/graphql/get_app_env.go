package graphql

import (
	"context"
	"fmt"
)

func (r *queryResolver) GetAppEnv(ctx context.Context, appName string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}
