package gopwn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestASM(t *testing.T) {
	t.Run("amd64", func(t *testing.T) {
		insn, err := AssembleAMD64("mov rax, 0")
		assert.NoError(t, err)
		assert.Equal(t, []byte("\x48\xc7\xc0\x00\x00\x00\x00"), insn)
	})

	t.Run("i386", func(t *testing.T) {
		insn, err := AssembleI386("mov eax, 0")
		assert.NoError(t, err)
		assert.Equal(t, []byte("\xb8\x00\x00\x00\x00"), insn)
	})
}

func TestDISASM(t *testing.T) {
	t.Run("amd64", func(t *testing.T) {
		assembly, err := DisamAMD64([]byte("\x48\xc7\xc0\x17\x00\x00\x00"), 0)
		assert.NoError(t, err)
		assert.Equal(t, "0x0           48 c7 c0 17 00 00 00          mov rax, 0x17", assembly)
	})

	t.Run("i386", func(t *testing.T) {
		assembly, err := DisamI386([]byte("\xb8\x5d\x00\x00\x00"), 0)
		assert.NoError(t, err)
		assert.Equal(t, "0x0           b8 5d 00 00 00                mov eax, 0x5d", assembly)
	})
}
