package singleton

import "sync"

type Singleton[T any] struct {
	entity      *T
	constructor func() *T
	lock        sync.Mutex
}

func NewSingleton[T any](constructor func() *T) Singleton[T] {
	var singleton Singleton[T]
	singleton.constructor = constructor
	return singleton
}

func (s *Singleton[T]) Get() *T {
	s.lock.Lock()
	if s.entity == nil {
		s.entity = s.constructor()
	}
	s.lock.Unlock()
	return s.entity
}
