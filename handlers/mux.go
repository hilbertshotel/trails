package handlers

import (
	"net/http"
	"trails/dep"
)

func Mux(d *dep.Dependencies) *http.ServeMux {
	mux := http.NewServeMux()

	// static files
	static := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	mux.Handle("/static/", static)

	// index
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index(w, r, d)
	})

	// send web socket address to client
	mux.HandleFunc("/wsa", func(w http.ResponseWriter, r *http.Request) {
		wsa(w, r, d)
	})

	// sorting handler
	mux.HandleFunc("/sorter", func(w http.ResponseWriter, r *http.Request) {
		sorter(w, r, d)
	})

	return mux
}
