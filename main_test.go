package main

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp/fasthttputil"
)

// serve serves http request using provided fasthttp handler
func serve(req *http.Request) (*http.Response, error) {
	ln := fasthttputil.NewInmemoryListener()
	defer ln.Close()

	app := newApp()

	errCh := make(chan error, 1)

	go func() {
		errCh <- app.Serve(ln)
	}()

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return ln.Dial()
			},
		},
	}

	return client.Do(req)
}

// $ go test -v
func TestNewApp(t *testing.T) {
	r, err := http.NewRequest("GET", "http://test/", nil)
	assert.Nil(t, err)

	res, err := serve(r)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(body))
}

func TestNotFound(t *testing.T) {
	r, err := http.NewRequest("GET", "http://test/notfound", nil)
	assert.Nil(t, err)

	res, err := serve(r)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.JSONEq(t, `{"status":404,"message":""}`, string(body))
}

func TestMethodNotAllowed(t *testing.T) {
	r, err := http.NewRequest("POST", "http://test/", nil)
	assert.Nil(t, err)

	res, err := serve(r)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.JSONEq(t, `{"status":405,"message":""}`, string(body))
}
