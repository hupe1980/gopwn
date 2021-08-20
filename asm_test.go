package gopwn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestASM(t *testing.T) {
	t.Run("x86_64", func(t *testing.T) {
		insn, err := Assemble_X86_64("mov rax, 0")
		assert.NoError(t, err)
		assert.Equal(t, []byte("\x48\xc7\xc0\x00\x00\x00\x00"), insn)
	})

	t.Run("i386", func(t *testing.T) {
		insn, err := Assemble_I386("mov eax, 0")
		assert.NoError(t, err)
		assert.Equal(t, []byte("\xb8\x00\x00\x00\x00"), insn)
	})
}

func TestDISASM(t *testing.T) {
	t.Run("x86_64", func(t *testing.T) {
		assembly, err := Disam_X86_64([]byte("\x48\xc7\xc0\x17\x00\x00\x00"), 0)
		assert.NoError(t, err)
		assert.Equal(t, "0x0:\tmov\t\trax, 0x17\n", assembly)
	})

	t.Run("i386", func(t *testing.T) {
		assembly, err := Disam_I386([]byte("\xb8\x5d\x00\x00\x00"), 0)
		assert.NoError(t, err)
		assert.Equal(t, "0x0:\tmov\t\teax, 0x5d\n", assembly)
	})
}
