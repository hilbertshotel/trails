package handlers

import (
	"net/http"
	"trails/dep"
)

func Mux(d *dep.Dependencies) *http.ServeMux {
	mux := http.NewServeMux()

	static := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	mux.Handle("/static/", static)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		get(w, r, d)
	})

	mux.HandleFunc("/workouts", func(w http.ResponseWriter, r *http.Request) {
		workouts(w, r, d)
	})

	return mux
}
