package singleton

import "sync"

type Singleton[T any] interface {
	Get() *T
}

type EagerSingleton[T any] struct {
	entity *T
}

func NewEagerSingleton[T any](constructor func() *T) EagerSingleton[T] {
	var singleton EagerSingleton[T]
	singleton.entity = constructor()
	return singleton
}

type LazySingleton[T any] struct {
	entity      *T
	constructor func() *T
	lock        sync.Mutex
}

func NewLazySingleton[T any](constructor func() *T) LazySingleton[T] {
	var singleton LazySingleton[T]
	singleton.constructor = constructor
	return singleton
}

func (s *LazySingleton[T]) Get() *T {
	s.lock.Lock()
	if s.entity == nil {
		s.entity = s.constructor()
	}
	s.lock.Unlock()
	return s.entity
}
