package gopwn

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHTTPGet(t *testing.T) {
	t.Run("default setup", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "helloworld")
		}))
		defer ts.Close()
		content, err := HTTPGet(ts.URL)
		assert.NoError(t, err)
		assert.Equal(t, []byte("helloworld"), content)
	})

	t.Run("timeout", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(10 * time.Millisecond)
		}))
		defer ts.Close()
		_, err := HTTPGet(ts.URL, func(o *HTTPClientOptions) {
			o.Timeout = 5 * time.Millisecond
		})
		assert.Error(t, err)
		netErr, ok := err.(net.Error)
		assert.True(t, ok)
		assert.True(t, netErr.Timeout())
	})
}
