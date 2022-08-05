package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
	"github.com/samber/lo"
)

// AppLogs is the resolver for the appLogs field.
func (r *subscriptionResolver) AppLogs(ctx context.Context, id string, excludeStdout *bool, excludeStderr *bool, initialCount *int) (<-chan *internal.Log, error) {
	service, err := r.RuntimeServiceRepo.GetOne(ctx, server.RuntimeServicesFilter{
		AppID: &id,
	})
	if err != nil {
		return nil, err
	}
	tail := lo.ToPtr("50")
	if initialCount != nil {
		*tail = fmt.Sprint(*initialCount)
	}
	stream, err := r.LogRepo.GetLogStream(ctx, server.LogsFilter{
		ServiceID:     service.ID,
		Follow:        lo.ToPtr(true),
		Tail:          tail,
		ExcludeStdout: excludeStdout,
		ExcludeStderr: excludeStderr,
	})
	if err != nil {
		return nil, err
	}

	logs := make(chan *internal.Log, 1)
	r.runSubscriptionJob(ctx, func(done func() bool) {
		defer stream.Close()
		for !done() {
			log, _, err := stream.NextLog()
			if err == nil {
				logs <- &log
			} else if !errors.Is(err, context.Canceled) {
				r.Logger.E("%v", err)
			}
		}
	})
	return logs, nil
}

// Subscription returns gqlgen.SubscriptionResolver implementation.
func (r *Resolver) Subscription() gqlgen.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
