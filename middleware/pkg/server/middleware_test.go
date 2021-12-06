package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCachingMiddleware(t *testing.T) {
	srv := NewServer(3)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})
	h := srv.CachingMiddleware(handler)

	if string(srv.Cache.Get("/hello")) != "" {
		t.Error("Cache should be empty, got: ", srv.Cache.Get("/hello"))
	}

	// Test that result is cached
	r := httptest.NewRequest(http.MethodGet, "/hello", nil)
	recorder := httptest.NewRecorder()
	h(recorder, r)
	if string(srv.Cache.Get("/hello")) != "Hello world!" {
		t.Error("Not cached properly, got: ", srv.Cache.Get("/hello"))
	}
}
