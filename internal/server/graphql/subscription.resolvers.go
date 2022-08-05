package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
)

// AppLog is the resolver for the appLog field.
func (r *subscriptionResolver) AppLog(ctx context.Context, id string) (<-chan *internal.Log, error) {
	service, err := r.RuntimeServiceRepo.GetOne(ctx, server.RuntimeServicesFilter{
		AppID: &id,
	})
	if err != nil {
		return nil, err
	}
	stream, err := r.LogRepo.GetLogStream(ctx, service.ID)
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
