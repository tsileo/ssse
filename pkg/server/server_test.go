package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"a4.io/ssse/pkg/client"
)

func TestServer(t *testing.T) {
	// New Server
	s := New()
	s.Start()

	mux := http.NewServeMux()
	mux.Handle("/events", s)
	server := httptest.NewServer(mux)

	url := server.URL + "/events"
	c := client.New(url)
	var cnt int
	go func() {
		if err := c.Subscribe(nil, func(e *client.Event) error {
			cnt++
			t.Logf("received event %+v\n", e)
			return nil
		}); err != nil {
			panic(err)
		}
	}()

	time.Sleep(1 * time.Second)
	for _, i := range []byte{1, 2, 3, 4, 5} {
		s.Publish("testing", []byte{i})
	}
	time.Sleep(1 * time.Second)

	if cnt != 5 {
		t.Errorf("failed, got %d events, expected 5", cnt)
	}
}
