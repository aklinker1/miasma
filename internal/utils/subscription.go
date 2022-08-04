package utils

import (
	"context"
	"math/rand"
	"sync"

	"github.com/aklinker1/miasma/internal/server"
	"github.com/samber/lo"
)

type SubscriptionManager[T any] struct {
	mu                  sync.Mutex
	name                string
	subscriptions       []*Subscription[T]
	logger              server.Logger
	activeSubscriptions int
	stopJob             func()
}

func NewSubscriptionManager[T any](logger server.Logger) *SubscriptionManager[T] {
	return &SubscriptionManager[T]{
		mu:                  sync.Mutex{},
		subscriptions:       []*Subscription[T]{},
		logger:              logger,
		activeSubscriptions: 0,
		stopJob:             func() {},
	}
}

func (m *SubscriptionManager[T]) CreateSubscription(job func(jobCtx context.Context, done *MutexValue[bool])) *Subscription[T] {
	sub := &Subscription[T]{
		id:      rand.Int63(),
		Channel: make(chan T, 1),
	}
	m.logger.V("Created %d", sub.id)

	m.mu.Lock()
	if m.activeSubscriptions == 0 {
		jobCtx, stopJob := context.WithCancel(context.Background())
		m.stopJob = stopJob
		go func() {
			done := NewMutexValue(false)
			go func() {
				<-jobCtx.Done()
				m.logger.V("Job canceled")
				done.Set(true)
			}()
			defer func() {
				m.logger.V("Job finished")
			}()
			job(jobCtx, done)
		}()
	}
	m.activeSubscriptions++
	m.subscriptions = append(m.subscriptions, sub)
	m.mu.Unlock()
	return sub
}

func (m *SubscriptionManager[T]) CancelSubscription(sub *Subscription[T]) {
	m.logger.V("Canceled %d", sub.id)

	m.mu.Lock()
	m.activeSubscriptions--
	if m.activeSubscriptions == 0 {
		m.stopJob()
	}
	m.subscriptions = lo.Filter(m.subscriptions, func(s *Subscription[T], _ int) bool {
		return s != sub
	})
	close(sub.Channel)
	m.mu.Unlock()
}

func (m *SubscriptionManager[T]) Broadcast(t T) {
	m.mu.Lock()
	m.logger.V("Broadcasting %v to %d subscriptions", t, len(m.subscriptions))
	// for _, s := range m.subscriptions {
	// 	s.Channel <- t
	// }
	m.mu.Unlock()
}

type Subscription[T any] struct {
	id      int64
	Channel chan T
}
