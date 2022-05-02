package graphql

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
)

func (r *mutationResolver) SetAppDockerConfig(ctx context.Context, appName string, newConfig *internal.DockerConfigInput) (*internal.DockerConfig, error) {
	panic(fmt.Errorf("not implemented"))
}
