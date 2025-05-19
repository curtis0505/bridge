package account

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	mand "math/rand"
	"strings"
	"sync/atomic"
	"time"
	// "net/url"
)

var objectIDCounter = readRandomUint32()
var processUnique = processUniqueBytes()

func RandStringRunes(n int) string {
	// var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()_+{}|[]")
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	mand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[mand.Intn(len(letterRunes))]
	}
	return string(b)
}

// NewObjectIDFromTimestamp generates a new ObjectID based on the given time.
// func NewID() []byte {
func NewGeneratorID() string {
	var b [16]byte

	binary.BigEndian.PutUint32(b[0:4], uint32(time.Now().UnixNano()))
	copy(b[4:9], processUnique[:])
	putUint24(b[9:12], atomic.AddUint32(&objectIDCounter, 1))
	copy(b[12:16], []byte(RandStringRunes(4)))

	res := fmt.Sprintf("%x", b)
	return res
}

func processUniqueBytes() [5]byte {
	var b [5]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(fmt.Errorf("cannot initialize objectid package with crypto.rand.Reader: %v", err))
	}

	return b
}

func readRandomUint32() uint32 {
	var b [4]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(fmt.Errorf("cannot initialize objectid package with crypto.rand.Reader: %v", err))
	}

	return (uint32(b[0]) << 0) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24)
}

func putUint24(b []byte, v uint32) {
	b[0] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[2] = byte(v)
}

func GenCodeCharset(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	var seededRand *mand.Rand = mand.New(mand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func getSourceIndexTotalByUUID(source string) int64 {

	// neopin user id인 uuid의 케릭터셋
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 62(index) * 32(source length) = (max)1984
	var totalInt64 int64

	for _, v := range source {
		totalInt64 += int64(strings.Index(charset, string(v))) + int64(1)
	}

	return totalInt64
}

func GenCodeCharsetBySource(source string, length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	sourceSeed := getSourceIndexTotalByUUID(source)

	var seededRand *mand.Rand = mand.New(mand.NewSource(time.Now().UnixNano() + sourceSeed))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
