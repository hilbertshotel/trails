package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"trails/dep"
	"trails/wrk"

	"go.mongodb.org/mongo-driver/bson"
)

func workouts(w http.ResponseWriter, r *http.Request, d *dep.Dependencies) {

	// handle method
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	// load workouts from db
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cursor, err := d.Coll.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}

	var workouts wrk.Workouts
	if err = cursor.All(ctx, &workouts); err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}

	// marshal workouts
	response, err := json.Marshal(workouts)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}

	// respond
	w.Header().Set("content-type", "application/json")
	w.Write(response)
}
