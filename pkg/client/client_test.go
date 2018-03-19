package client

import "testing"

// This is the only test, the client is used and tested in ../server/server_test.go
func TestInvalidClient(t *testing.T) {
	c := New("http://google.com")
	if err := c.Subscribe(nil, nil); err != ErrInvalidSubscribeArgs {
		t.Errorf("client should have returned %v, got %v", ErrInvalidSubscribeArgs, err)
	}
}
