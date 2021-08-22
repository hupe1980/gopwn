package gopwn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestELF(t *testing.T) {
	t.Run("Open i386", func(t *testing.T) {
		elf, err := NewELF("testdata/elf.i386")
		assert.NoError(t, err)
		defer elf.Close()

		assert.Equal(t, ARCH_I386.String(), elf.Architecture().String())
	})

	t.Run("Open x86_64", func(t *testing.T) {
		elf, err := NewELF("testdata/elf.x86_64")
		assert.NoError(t, err)
		defer elf.Close()

		assert.Equal(t, ARCH_X86_64.String(), elf.Architecture().String())
	})

	t.Run("Patch i386", func(t *testing.T) {
		elf, err := NewELF("testdata/elf.i386")
		assert.NoError(t, err)
		defer elf.Close()

		err = elf.Write([]byte{0x41, 0x41, 0x41, 0x41}, 0x1090)
		assert.NoError(t, err)

		b, err := elf.Read(0x1090, 4)
		assert.NoError(t, err)
		assert.Equal(t, []byte{0x41, 0x41, 0x41, 0x41}, b)
	})
}
