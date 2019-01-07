package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/gorilla/mux"
)

// thanks to https://gist.github.com/renthraysk/e3a8f68252d120ffce1442fb02fdb34a
func routeLog(r *http.Request) *log.Entry {
	l := log.WithFields(log.Fields{
		"method":     r.Method,
		"requestURI": r.RequestURI,
	})
	return l
}

func index(w http.ResponseWriter, r *http.Request) {
	log := routeLog(r)
	log.Info("index")
	fmt.Fprintf(w, "hello")
}

func about(w http.ResponseWriter, r *http.Request) {
	log := routeLog(r)
	log.Info("about")
	fmt.Fprintf(w, "about")
}

func main() {
	app := mux.NewRouter()
	app.HandleFunc("/", index)
	app.HandleFunc("/about", about)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), app); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}
