package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/gorilla/mux"
)

type key int

const (
	logger key = iota
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging := log.WithFields(log.Fields{
			"method":     r.Method,
			"requestURI": r.RequestURI,
		})
		ctx := context.WithValue(r.Context(), logger, logging)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	log, ok := r.Context().Value(logger).(*log.Entry)
	if !ok {
		http.Error(w, "Unable to get logging context", http.StatusInternalServerError)
		return
	}
	log.Info("index")
	fmt.Fprintf(w, "hello")
}

func about(w http.ResponseWriter, r *http.Request) {
	log, ok := r.Context().Value(logger).(*log.Entry)
	if !ok {
		http.Error(w, "Unable to get logging context", http.StatusInternalServerError)
		return
	}
	log.Info("about")
	fmt.Fprintf(w, "about")
}

func main() {
	app := mux.NewRouter()
	app.HandleFunc("/", index)
	app.HandleFunc("/about", about)
	app.Use(loggingMiddleware)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), app); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}
