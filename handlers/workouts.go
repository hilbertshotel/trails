package handlers

import (
	"encoding/json"
	"net/http"
	"trails/dep"
	"trails/wrk"
)

func workouts(w http.ResponseWriter, r *http.Request, d *dep.Dependencies) {

	// Handle method
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	// Load workouts from db
	workouts, err := wrk.Load(d.DB)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}

	// Turn duration to seconds for easier frontend sorting
	err = (&workouts).ParseDuration()
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}

	// Marshal data
	response, err := json.Marshal(workouts)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}

	// Return data to frontend
	w.Header().Set("content-type", "application/json")
	w.Write(response)
}
