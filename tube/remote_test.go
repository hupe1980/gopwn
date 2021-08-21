package tube

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRecvLine(t *testing.T) {
	go func() {
		l, err := net.Listen("tcp", ":8000")
		if err != nil {
			t.Error(err)
		}
		defer l.Close()

		conn, err := l.Accept()
		if err != nil {
			t.Error(err)
		}
		defer conn.Close()

		if _, err := conn.Write([]byte("helloworld\n")); err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(500 * time.Millisecond)
	r, err := NewRemote("tcp", ":8000")
	assert.NoError(t, err)
	defer r.Close()

	data, err := r.RecvLine()
	assert.NoError(t, err)
	assert.Equal(t, []byte("helloworld"), data)
}
