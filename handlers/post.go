package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"trails/logger"
	"trails/models"
)

func _(w http.ResponseWriter, r *http.Request, log *logger.Logger) {

	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

	var workout models.Workout
	err = json.Unmarshal(data, &workout)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

	// insert session into database
}
