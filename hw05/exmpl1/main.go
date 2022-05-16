package main

import (
	"fmt"
	"log"
	"net/http"
)

type MuxHandler struct {
	GetHandler  http.Handler
	PostHandler http.Handler
}

func (h MuxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetHandler.ServeHTTP(w, r)
	case http.MethodPost:
		h.PostHandler.ServeHTTP(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := MuxHandler{
		GetHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "GET HANDLER")
		}),
		PostHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "POST HANDLER")
		}),
	}
	http.Handle("/", mux)
	log.Fatal(http.ListenAndServe(":8020", nil))
}
