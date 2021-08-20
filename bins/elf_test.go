package bins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestELF(t *testing.T) {
	t.Run("i386", func(t *testing.T) {
		elf, err := NewELF("../testdata/elf.i386")
		assert.NoError(t, err)
		defer elf.Close()

		assert.Equal(t, ARCH_I386.String(), elf.Architecture().String())
	})

	t.Run("x86_64", func(t *testing.T) {
		elf, err := NewELF("../testdata/elf.x86_64")
		assert.NoError(t, err)
		defer elf.Close()

		assert.Equal(t, ARCH_X86_64.String(), elf.Architecture().String())
	})
}
