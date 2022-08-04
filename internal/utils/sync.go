package utils

import "sync"

type MutexValue[T any] struct {
	value T
	mu    sync.Mutex
}

func NewMutexValue[T any](t T) *MutexValue[T] {
	return &MutexValue[T]{
		value: t,
		mu:    sync.Mutex{},
	}
}

func (v *MutexValue[T]) Value() T {
	v.mu.Lock()
	t := v.value
	v.mu.Unlock()
	return t
}

func (v *MutexValue[T]) Set(t T) {
	v.mu.Lock()
	v.value = t
	v.mu.Unlock()
}

func (v *MutexValue[T]) Do(fn func(t T)) {
	v.mu.Lock()
	fn(v.value)
	v.mu.Unlock()
}

func (v *MutexValue[T]) Mutate(fn func(t T) T) {
	v.mu.Lock()
	v.value = fn(v.value)
	v.mu.Unlock()
}
