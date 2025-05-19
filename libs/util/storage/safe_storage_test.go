package storage

import (
	mongoServiceDB "github.com/curtis0505/bridge/libs/database/mongo/service_db"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSafeStorageSimpleLoadAndStore(t *testing.T) {
	// TEST: store
	storage := NewSafeStorage[string, *mongoServiceDB.Tokens]()
	storage.Store("test1", &mongoServiceDB.Tokens{CurrencyID: "ETH001"})
	storage.Store("test2", &mongoServiceDB.Tokens{CurrencyID: "ETH002"})
	storage.Store("test3", &mongoServiceDB.Tokens{CurrencyID: "ETH003"})

	// TEST: load
	tokenInfo1, ok := storage.Load("test1")
	assert.Equal(t, true, ok)
	assert.Equal(t, "ETH001", tokenInfo1.CurrencyID)

	tokenInfo2, ok := storage.Load("test2")
	assert.Equal(t, true, ok)
	assert.Equal(t, "ETH002", tokenInfo2.CurrencyID)

	tokenInfo3, ok := storage.Load("test3")
	assert.Equal(t, true, ok)
	assert.Equal(t, "ETH003", tokenInfo3.CurrencyID)
}

func TestSafeStorageUsingPointer(t *testing.T) {
	storage := NewSafeStorage[string, *mongoServiceDB.Tokens]()
	tokenInfo1 := &mongoServiceDB.Tokens{CurrencyID: "ETH001"}
	storage.Store("test1", tokenInfo1)

	// TEST: change value
	tokenInfo1.CurrencyID = "WETH001"

	// TEST: load
	tmp, ok := storage.Load("test1")
	assert.Equal(t, true, ok)
	assert.Equal(t, "ETH001", tmp.CurrencyID)
}

func BenchmarkSafeStorage_Store(b *testing.B) {
	storage := NewSafeStorage[string, *mongoServiceDB.Tokens]()
	for i := 0; i < b.N; i++ {
		storage.Store("test", &mongoServiceDB.Tokens{CurrencyID: "ETH001"})
	}
}

func BenchmarkSafeStorage_Load(b *testing.B) {
	storage := NewSafeStorage[string, *mongoServiceDB.Tokens]()
	storage.Store("test", &mongoServiceDB.Tokens{CurrencyID: "ETH001"})
	for i := 0; i < b.N; i++ {
		_, _ = storage.Load("test")
	}
}

func BenchmarkSafeStorage_ForEach(b *testing.B) {
	storage := NewSafeStorage[string, *mongoServiceDB.Tokens]()
	for i := 0; i < 1000; i++ {
		storage.Store("test"+strconv.Itoa(i), &mongoServiceDB.Tokens{CurrencyID: "ETH001"})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		storage.ForEach(func(k string, v *mongoServiceDB.Tokens) error {
			return nil
		})
	}
}
