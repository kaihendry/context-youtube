package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gorilla/mux"
)

func TestRace(t *testing.T) {
	h := New()
	app := mux.NewRouter()
	app.HandleFunc("/", h.index)
	app.HandleFunc("/about", h.about)
	app.Use(h.loggingMiddleware)

	server := httptest.NewServer(app)
	defer server.Close()

	wg := &sync.WaitGroup{}
	for _, path := range []string{"/", "/about"} {
		path := path
		wg.Add(1)
		go func() {
			defer wg.Done()
			req, err := http.NewRequest(http.MethodGet, server.URL+path, nil)
			fmt.Println(server.URL + path)
			if err != nil {
				panic(err)
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
		}()
	}
	wg.Wait()
}
