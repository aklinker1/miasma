package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
	"github.com/aklinker1/miasma/internal/utils"
)

// AppLog is the resolver for the appLog field.
func (r *subscriptionResolver) AppLog(ctx context.Context, id string) (<-chan []*internal.Log, error) {
	service, err := r.RuntimeServiceRepo.GetOne(ctx, server.RuntimeServicesFilter{AppID: &id})
	if err != nil {
		return nil, err
	}

	sub := r.LogSubscriptions.CreateSubscription(func(jobCtx context.Context, done *utils.MutexValue[bool]) {
		stream, err := r.LogRepo.GetLogStream(jobCtx, service.ID)
		if err != nil {

			return
		}
		defer stream.Close()

		for !done.Value() {
			r.Logger.D("Waiting for log...")
			log, noMoreLogs, err := stream.NextLog()
			if noMoreLogs {
				r.Logger.I("No more logs")
				break
			} else if err != nil {
				r.Logger.E("Failed to read line: %v", err)
				break
			}
			r.Logger.V("[%v] %s", log.Timestamp, log.Message)
			r.LogSubscriptions.Broadcast([]*internal.Log{&log})
		}

	})

	go func() {
		<-ctx.Done()
		r.Logger.V("Websocket closed")
		r.LogSubscriptions.CancelSubscription(sub)
	}()

	// go func() {
	// 	r.Logger.I("Watcher started")
	// 	for true {
	// 		r.Logger.D("Waiting for log...")
	// 		log, noMoreLogs, err := stream.NextLog()
	// 		if noMoreLogs {
	// 			r.Logger.I("No more logs")
	// 			break
	// 		} else if err != nil {
	// 			r.Logger.E("Failed to read line: %v", err)
	// 			break
	// 		}
	// 		r.Logger.V("[%v] %s", log.Timestamp, log.Message)
	// 		sub <- []*internal.Log{&log}
	// 	}
	// 	done()
	// 	r.Logger.I("Watcher done")
	// }()

	sub.Channel <- []*internal.Log{}

	// done()
	return sub.Channel, nil
}

// Subscription returns gqlgen.SubscriptionResolver implementation.
func (r *Resolver) Subscription() gqlgen.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
