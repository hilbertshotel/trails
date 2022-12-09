package handlers

import (
	"context"
	"net/http"
	"time"
	"trails/dep"
	"trails/wrk"

	"go.mongodb.org/mongo-driver/bson"
)

func get(w http.ResponseWriter, r *http.Request, d *dep.Dependencies) {

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

	// get totals
	total, err := workouts.ParseTotal(d.Log)
	if err != nil {
		d.Log.Error(err)
		return
	}

	// get best workout
	best := workouts.FindBest()

	// gather data
	data := wrk.Data{
		Total:    total,
		Best:     best,
		Workouts: workouts,
	}

	// return template
	if err := d.Tmp.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}

}
