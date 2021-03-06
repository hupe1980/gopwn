package gopwn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5Sum(t *testing.T) {
	hash := MD5Sum([]byte("ABCD"))
	assert.Equal(t, "cb08ca4a7bb5f9683c19133a84872ca7", hash)
}

func TestHex(t *testing.T) {
	b := Hex([]byte("ABCD"))
	assert.Equal(t, []byte("\x34\x31\x34\x32\x34\x33\x34\x34"), b)
}

func TestHexString(t *testing.T) {
	s := HexString([]byte("ABCD"))
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

func TestBase64E(t *testing.T) {
	s := Base64E([]byte("ABCD"))
	assert.Equal(t, "QUJDRA==", s)
}

func TestBase64D(t *testing.T) {
	b, err := Base64D("QUJDRA==")
	assert.NoError(t, err)
	assert.Equal(t, []byte("ABCD"), b)
}

func TestROT13(t *testing.T) {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	assert.Equal(t, s, ROT13(ROT13(s)))
}

func TestOpenFile(t *testing.T) {
	t.Run("ELF I386", func(t *testing.T) {
		fh, bt, err := OpenFile("testdata/elf.i386")
		assert.NoError(t, err)
		defer fh.Close()
		assert.Equal(t, BINTYPE_ELF, bt)
	})

	t.Run("ELF AMD64", func(t *testing.T) {
		fh, bt, err := OpenFile("testdata/elf.amd64")
		assert.NoError(t, err)
		defer fh.Close()
		assert.Equal(t, BINTYPE_ELF, bt)
	})

	t.Run("MACHO", func(t *testing.T) {
		fh, bt, err := OpenFile("testdata/macho.amd64")
		assert.NoError(t, err)
		defer fh.Close()
		assert.Equal(t, BINTYPE_MACHO, bt)
	})
}
