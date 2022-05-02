package graphql

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
)

func (r *queryResolver) GetAppDockerConfig(ctx context.Context, appName string) (*internal.DockerConfig, error) {
	panic(fmt.Errorf("not implemented"))
}
