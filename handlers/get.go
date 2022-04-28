package handlers

import (
	"database/sql"
	"net/http"
	t "text/template"
	l "trails/logger"
	"trails/models"
)

func get(w http.ResponseWriter, r *http.Request, log *l.Logger, tmp *t.Template, db *sql.DB) {

	// handle method
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	// load workouts from db
	workouts, err := models.QueryWorkouts(db)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

	// analyze workouts
	data, err := models.Analyize(workouts)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

	// return template
	if err := tmp.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

}
