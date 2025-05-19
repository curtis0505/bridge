package neo_aes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CBCDecrypt(t *testing.T) {
	privKey := []byte("09coycgxq9zj0mx9r3f8yaub0n71gmzy")
	password := "bc91a5946e7a4f5c992fc8cf1c5d88cf:798ed32ab4e3644612b94edc5ec9cbdb"

	decrypt, err := CBCDecrypt(privKey, password)
	assert.NoError(t, err, "Error should be nil")
	assert.NotNil(t, decrypt, "Decrypted text should not be nil")
	assert.NotEmpty(t, decrypt, "Decrypted text should not be empty")

	expected := "npdev12#$" // Replace with the actual expected decrypted text
	assert.Equal(t, expected, string(decrypt), "Decrypted text should match the expected value")
}
