package graphql

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
)

func (r *queryResolver) GetPlugin(ctx context.Context, pluginName string) (*internal.Plugin, error) {
	panic(fmt.Errorf("not implemented"))
}
