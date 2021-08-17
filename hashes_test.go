package gopwn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5Sum(t *testing.T) {
	hash := MD5Sum([]byte("ABCD"))
	assert.Equal(t, "cb08ca4a7bb5f9683c19133a84872ca7", hash)
}
