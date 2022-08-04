package utils

import (
	"math/rand"
	"sync"
)

type SubscriptionManager[T any] struct {
	subscriptions sync.Map
}

func NewSubscriptionManager[T any]() *SubscriptionManager[T] {
	return &SubscriptionManager[T]{
		subscriptions: sync.Map{},
	}
}

func (m *SubscriptionManager[T]) CreateSubscription() *Subscription[T] {
	sub := &Subscription[T]{
		id:      rand.Int63(),
		Channel: make(chan T, 1),
	}
	m.subscriptions.Store(sub.id, sub)
	return sub
}

func (m *SubscriptionManager[T]) CancelSubscription(sub *Subscription[T]) {
	close(sub.Channel)
	m.subscriptions.Delete(sub.id)
}

type Subscription[T any] struct {
	id      int64
	Channel chan T
}
