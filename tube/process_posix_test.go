// +build linux freebsd openbsd darwin solaris

package tube

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	t.Run("echo", func(t *testing.T) {
		p, err := NewProcess([]string{"echo", "helloworld"})
		assert.NoError(t, err)

		out, err := p.RecvLine()
		assert.NoError(t, err)
		assert.Equal(t, []byte("helloworld"), out)
	})

	t.Run("sh", func(t *testing.T) {
		p, err := NewProcess([]string{"sh"})
		assert.NoError(t, err)

		_, err = p.SendLine("echo helloworld")
		assert.NoError(t, err)

		out, err := p.RecvLine()
		assert.NoError(t, err)
		assert.Equal(t, []byte("helloworld"), out)
	})

	t.Run("env", func(t *testing.T) {
		p, err := NewProcess([]string{"sh"}, func(o *ProcessOptions) {
			o.Env = []string{"HELLO_WORLD=helloworld"}
		})
		assert.NoError(t, err)

		_, err = p.SendLine("echo $HELLO_WORLD")
		assert.NoError(t, err)

		out, err := p.RecvLine()
		assert.NoError(t, err)
		assert.Equal(t, []byte("helloworld"), out)
	})

	t.Run("cwd", func(t *testing.T) {
		dir, _ := filepath.Abs("../testdata")
		p, err := NewProcess([]string{"sh"}, func(o *ProcessOptions) {
			o.Dir = dir
		})
		assert.NoError(t, err)

		_, err = p.SendLine("cd .. && pwd")
		assert.NoError(t, err)

		out, err := p.RecvLine()
		assert.NoError(t, err)
		assert.Equal(t, "gopwn", filepath.Base(string(out)))

		cwd, err := p.CWD()
		assert.NoError(t, err)
		assert.Equal(t, "gopwn", filepath.Base(cwd))
	})
}
