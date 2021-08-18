package tubes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	t.Run("echo", func(t *testing.T) {
		p, err := NewProcess([]string{"echo", "helloworld"})
		assert.NoError(t, err)

		err = p.Start()
		assert.NoError(t, err)

		out, err := p.RecvLine()
		assert.NoError(t, err)
		assert.Equal(t, []byte("helloworld\n"), out)
	})

	t.Run("sh", func(t *testing.T) {
		p, err := NewProcess([]string{"sh"})
		assert.NoError(t, err)

		err = p.Start()
		assert.NoError(t, err)

		_, err = p.SendLine("echo helloworld")
		assert.NoError(t, err)

		out, err := p.RecvLine()
		assert.NoError(t, err)
		assert.Equal(t, []byte("helloworld\n"), out)
	})

	t.Run("env", func(t *testing.T) {
		p, err := NewProcess([]string{"sh"}, func(o *ProcessOptions) {
			o.Env = []string{"HELLO_WORLD=helloworld"}
		})
		assert.NoError(t, err)

		err = p.Start()
		assert.NoError(t, err)

		_, err = p.SendLine("echo $HELLO_WORLD")
		assert.NoError(t, err)

		out, err := p.RecvLine()
		assert.NoError(t, err)
		assert.Equal(t, []byte("helloworld\n"), out)
	})
}
