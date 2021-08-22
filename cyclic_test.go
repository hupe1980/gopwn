package gopwn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCycling(t *testing.T) {
	t.Run("default setup", func(t *testing.T) {
		l := 16
		pattern := Cyclic(l)
		assert.Equal(t, l, len(pattern))
		assert.Equal(t, []byte("aaaabaaacaaadaaa"), pattern)
	})

	t.Run("custom alphabet", func(t *testing.T) {
		l := 16
		pattern := Cyclic(l, func(o *CyclicOptions) {
			o.Alphabet = "XYZ"
		})
		assert.Equal(t, l, len(pattern))
		assert.Equal(t, []byte("XXXXYXXXZXXYYXXY"), pattern)
	})

	t.Run("custom n", func(t *testing.T) {
		l := 16
		pattern := Cyclic(l, func(o *CyclicOptions) {
			o.DistSubseqLength = 2
		})
		assert.Equal(t, l, len(pattern))
		assert.Equal(t, []byte("aabacadaeafagaba"), pattern)
	})
}

func TestCyclingFind(t *testing.T) {
	offset := CyclicFind([]byte("baaa"))
	assert.Equal(t, 4, offset)
}
