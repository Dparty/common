package singleton

import "sync"

type Singleton[T any] interface {
	Get() *T
}

type EagerSingleton[T any] struct {
	instance *T
}

func (s *EagerSingleton[T]) Get() *T {
	return s.instance
}

func NewSingleton[T any](constructor func() *T, eagerModel bool) Singleton[T] {
	if eagerModel {
		return NewEagerSingleton[T](constructor)
	} else {
		return NewLazySingleton[T](constructor)
	}
}

func NewEagerSingleton[T any](constructor func() *T) *EagerSingleton[T] {
	var singleton EagerSingleton[T]
	singleton.instance = constructor()
	return &singleton
}

type LazySingleton[T any] struct {
	instance    *T
	constructor func() *T
	lock        sync.Mutex
}

func NewLazySingleton[T any](constructor func() *T) *LazySingleton[T] {
	var singleton LazySingleton[T]
	singleton.constructor = constructor
	return &singleton
}

func (s *LazySingleton[T]) Get() *T {
	s.lock.Lock()
	if s.instance == nil {
		s.instance = s.constructor()
	}
	s.lock.Unlock()
	return s.instance
}
