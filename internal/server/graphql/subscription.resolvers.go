package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
)

// AppLog is the resolver for the appLog field.
func (r *subscriptionResolver) AppLog(ctx context.Context, id string) (<-chan *internal.Log, error) {
	// logs := make(chan *internal.Log, 1)
	// runSubscriptionJob(ctx, func(done func() bool) {
	// 	for i := 1; !done(); i++ {
	// 		message := fmt.Sprintf("Log %d", i)
	// 		fmt.Println("Sending", message)
	// 		logs <- &internal.Log{
	// 			Message:   message,
	// 			Timestamp: time.Now(),
	// 		}
	// 		time.Sleep(1 * time.Second)
	// 	}
	// })
	// return logs, nil

	_, err := r.RuntimeServiceRepo.GetOne(ctx, server.RuntimeServicesFilter{
		AppID: &id,
	})
	if err != nil {
		return nil, err
	}

	// stream, err := r.LogRepo.GetLogStream(ctx, service.ID)
	// if err != nil {
	// 	return nil, err
	// }
	// defer stream.Close()

	logs := make(chan *internal.Log, 1)
	r.runSubscriptionJob(ctx, func(done func() bool) {
		for i := 1; !done(); i++ {
			message := fmt.Sprintf("Log %d", i)
			r.Logger.V("Sending %s", message)
			logs <- &internal.Log{
				Message:   message,
				Timestamp: time.Now(),
			}
			time.Sleep(1 * time.Second)
		}
	})
	return logs, nil
}

// Subscription returns gqlgen.SubscriptionResolver implementation.
func (r *Resolver) Subscription() gqlgen.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
