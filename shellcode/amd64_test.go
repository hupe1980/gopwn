package shellcode

import (
	"strings"
	"testing"

	"github.com/hupe1980/gopwn"
	"github.com/stretchr/testify/assert"
)

func TestAMD64(t *testing.T) {
	amd64, err := NewAMD64()
	assert.NoError(t, err)

	t.Run("NOP", func(t *testing.T) {
		assert.Equal(t, "nop", strings.TrimSpace(amd64.NOP()))
	})

	t.Run("Linux", func(t *testing.T) {
		t.Run("SH", func(t *testing.T) {
			b, err := gopwn.AssembleAMD64(amd64.Linux.SH())
			assert.NoError(t, err)
			assert.Equal(t, []byte("\x31\xc0\x48\xbb\xd1\x9d\x96\x91\xd0\x8c\x97\xff\x48\xf7\xdb\x53\x54\x5f\x99\x52\x57\x54\x5e\xb0\x3b\x0f\x05"), b)
		})
	})
}
