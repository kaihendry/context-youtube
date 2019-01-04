package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Infof("%s - %s", r.Method, r.RequestURI)
	fmt.Fprintf(w, "hello")
}

func main() {
	app := mux.NewRouter()
	app.PathPrefix("/").HandlerFunc(index)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), app); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}
