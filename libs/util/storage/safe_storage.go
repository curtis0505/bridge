package storage

import (
	"reflect"
	"sync"
)

type Cloneable[V any] interface {
	Clone() V
}

type safeStorage[K any, V Cloneable[V]] struct {
	mu sync.RWMutex
	m  *sync.Map
}

func (s *safeStorage[K, V]) Load(k K) (V, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.m.Load(k)
	if !ok {
		var zeroV V
		return zeroV, false
	}
	return v.(V).Clone(), ok
}

func (s *safeStorage[K, V]) Store(k K, v V) {
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return
		}
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.m.Store(k, v.Clone())
}

func (s *safeStorage[K, V]) Delete(k K) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m.Delete(k)
}

func (s *safeStorage[K, V]) ForEach(f func(K, V) error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.m.Range(func(k, v any) bool {
		return f(k.(K), v.(V).Clone()) == nil
	})
}

func (s *safeStorage[K, V]) LoadOrStore(k K, v V) (V, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	actual, loaded := s.m.LoadOrStore(k, v)
	if loaded {
		return actual.(V).Clone(), loaded
	}

	return v.Clone(), loaded
}

func (s *safeStorage[K, V]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m = &sync.Map{}
}

func (s *safeStorage[K, V]) Replace(storage Storage[K, V]) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m = storage.(*safeStorage[K, V]).m
}

func NewSafeStorage[K any, V Cloneable[V]]() Storage[K, V] {
	return &safeStorage[K, V]{
		m:  &sync.Map{},
		mu: sync.RWMutex{},
	}
}
