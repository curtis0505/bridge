package storage

import (
	mongoServiceDB "github.com/curtis0505/bridge/libs/database/mongo/service_db"
	"github.com/holiman/uint256"
	"github.com/shopspring/decimal"
	"reflect"
	"sync"
)

var (
	_ Storage[int, string]                    = &syncMapStorage[int, string]{}
	_ Storage[string, float64]                = &syncMapStorage[string, float64]{}
	_ Storage[string, *uint256.Int]           = &syncMapStorage[string, *uint256.Int]{}
	_ Storage[string, decimal.Decimal]        = &syncMapStorage[string, decimal.Decimal]{}
	_ Storage[string, *mongoServiceDB.Tokens] = &syncMapStorage[string, *mongoServiceDB.Tokens]{}
)

type Storage[K any, V any] interface {
	Load(K) (V, bool)
	Store(K, V)
	Delete(K)
	ForEach(func(K, V) (err error))
	LoadOrStore(K, V) (V, bool)
	Clear()
	Replace(Storage[K, V])
}

type syncMapStorage[K any, V any] struct {
	mu sync.RWMutex
	m  *sync.Map
}

func (s *syncMapStorage[K, V]) Load(k K) (V, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.m.Load(k)
	if !ok {
		var zeroV V
		return zeroV, false
	}
	return v.(V), ok
}

func (s *syncMapStorage[K, V]) Store(k K, v V) {
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return
		}
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.m.Store(k, v)
}

func (s *syncMapStorage[K, V]) Delete(k K) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m.Delete(k)
}

func (s *syncMapStorage[K, V]) ForEach(f func(K, V) error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.m.Range(func(k, v any) bool {
		return f(k.(K), v.(V)) == nil
	})
}

func (s *syncMapStorage[K, V]) LoadOrStore(k K, v V) (V, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	actual, loaded := s.m.LoadOrStore(k, v)
	if loaded {
		return actual.(V), loaded
	}

	return v, loaded
}

func (s *syncMapStorage[K, V]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m = &sync.Map{}
}

func (s *syncMapStorage[K, V]) Replace(storage Storage[K, V]) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m = storage.(*syncMapStorage[K, V]).m
}

func New[K any, V any]() Storage[K, V] {
	return &syncMapStorage[K, V]{
		m:  &sync.Map{},
		mu: sync.RWMutex{},
	}
}
