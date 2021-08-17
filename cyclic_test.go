package gopwn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCycling(t *testing.T) {
	l := 16
	pattern := Cyclic(l)
	assert.Equal(t, l, len(pattern))
	assert.Equal(t, "aaaabaaacaaadaaa", pattern)
}

func TestCyclingFind(t *testing.T) {
	offset := CyclicFind([]byte("baaa"))
	assert.Equal(t, 4, offset)
}
