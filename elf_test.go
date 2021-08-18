package gopwn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestELF(t *testing.T) {
	elf, err := NewELF("testdata/cat")
	assert.NoError(t, err)
	defer elf.Close()
}
