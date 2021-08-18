package gopwn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPack(t *testing.T) {
	t.Run("16 LittleEndian", func(t *testing.T) {
		b := P16L(0xdead)
		assert.Equal(t, []byte("\xad\xde"), b)
	})

	t.Run("16 BigEndian", func(t *testing.T) {
		b := P16B(0xdead)
		assert.Equal(t, []byte("\xde\xad"), b)
	})

	t.Run("32 LittleEndian", func(t *testing.T) {
		b := P32L(0xdeadbeef)
		assert.Equal(t, []byte("\xef\xbe\xad\xde"), b)
	})

	t.Run("32 BigEndian", func(t *testing.T) {
		b := P32B(0xdeadbeef)
		assert.Equal(t, []byte("\xde\xad\xbe\xef"), b)
	})

	t.Run("64 LittleEndian", func(t *testing.T) {
		b := P64L(0xdeadbeef)
		assert.Equal(t, []byte("\xef\xbe\xad\xde\x00\x00\x00\x00"), b)
	})

	t.Run("64 BigEndian", func(t *testing.T) {
		b := P64B(0xdeadbeef)
		assert.Equal(t, []byte("\x00\x00\x00\x00\xde\xad\xbe\xef"), b)
	})
}

func TestUnPack(t *testing.T) {
	t.Run("16 LittleEndian", func(t *testing.T) {
		i := U16L([]byte("\xad\xde"))
		assert.Equal(t, uint16(0xdead), i)
	})

	t.Run("16 BigEndian", func(t *testing.T) {
		i := U16B([]byte("\xde\xad"))
		assert.Equal(t, uint16(0xdead), i)
	})

	t.Run("32 LittleEndian", func(t *testing.T) {
		i := U32L([]byte("\xef\xbe\xad\xde"))
		assert.Equal(t, uint32(0xdeadbeef), i)
	})

	t.Run("32 BigEndian", func(t *testing.T) {
		i := U32B([]byte("\xde\xad\xbe\xef"))
		assert.Equal(t, uint32(0xdeadbeef), i)
	})

	t.Run("64 LittleEndian", func(t *testing.T) {
		i := U64L([]byte("\xef\xbe\xad\xde\x00\x00\x00\x00"))
		assert.Equal(t, uint64(0xdeadbeef), i)
	})

	t.Run("64 BigEndian", func(t *testing.T) {
		i := U64B([]byte("\x00\x00\x00\x00\xde\xad\xbe\xef"))
		assert.Equal(t, uint64(0xdeadbeef), i)
	})
}
