package ui

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	log.Println("Server is starting")
	r := mux.NewRouter()
	r.HandleFunc("/", staticFileHandler)
	r.HandleFunc("/api/proxy", proxyUpdaterHandler)
	http.Handle("/", r)
}

func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Static files -- stub in later")
}

func proxyUpdaterHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/api/proxy was hit")
}
