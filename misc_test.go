package gopwn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHex(t *testing.T) {
	b := Hex([]byte("ABCD"))
	assert.Equal(t, []byte("\x34\x31\x34\x32\x34\x33\x34\x34"), b)
}

func TestHexToString(t *testing.T) {
	s := HexToString([]byte("ABCD"))
	assert.Equal(t, "41424344", s)
}

func TestUnHex(t *testing.T) {
	b, err := UnHex([]byte("\x34\x31\x34\x32\x34\x33\x34\x34"))
	assert.NoError(t, err)
	assert.Equal(t, []byte("ABCD"), b)
}

func TestUnhexString(t *testing.T) {
	b, err := UnHexString("41424344")
	assert.NoError(t, err)
	assert.Equal(t, []byte("ABCD"), b)
}
