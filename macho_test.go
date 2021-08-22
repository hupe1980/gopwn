package gopwn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMACHO(t *testing.T) {
	macho, err := NewMACHO("testdata/macho.amd64")
	assert.NoError(t, err)
	defer macho.Close()
}
