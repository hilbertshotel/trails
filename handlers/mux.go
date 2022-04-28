package handlers

import (
	"net/http"
	"text/template"
	"trails/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

func Mux(log *logger.Logger, tmp *template.Template, coll *mongo.Collection) *http.ServeMux {
	mux := http.NewServeMux()

	static := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	mux.Handle("/static/", static)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		get(w, r, log, tmp, coll)
	})

	return mux
}
