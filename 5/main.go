package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/gorilla/mux"
)

type handler struct{ Log *log.Entry }

func (h *handler) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Log = log.WithFields(log.Fields{
			"method":     r.Method,
			"requestURI": r.RequestURI,
		})
		next.ServeHTTP(w, r)
	})
}

func New() (h *handler) { return &handler{Log: log.WithFields(log.Fields{"test": "FAIL"})} }

func (h *handler) index(w http.ResponseWriter, r *http.Request) {
	// log.Infof("%s - %s", r.Method, r.RequestURI)
	h.Log.Info("index")
	fmt.Fprintf(w, "hello")
}

func (h *handler) about(w http.ResponseWriter, r *http.Request) {
	// log.Infof("%s - %s", r.Method, r.RequestURI)
	h.Log.Info("about")
	fmt.Fprintf(w, "about")
}

func main() {
	h := New()
	app := mux.NewRouter()
	app.HandleFunc("/", h.index)
	app.HandleFunc("/about", h.about)
	app.Use(h.loggingMiddleware)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), app); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}
