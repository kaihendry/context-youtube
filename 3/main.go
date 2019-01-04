package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"requestURI": r.RequestURI,
		}).Info("Middleware")
		next.ServeHTTP(w, r)
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	// log.Infof("%s - %s", r.Method, r.RequestURI)
	fmt.Fprintf(w, "hello")
}

func about(w http.ResponseWriter, r *http.Request) {
	// log.Infof("%s - %s", r.Method, r.RequestURI)
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
