package handlers

import (
	"context"
	"net/http"
	t "text/template"
	"time"
	l "trails/logger"
	"trails/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func get(w http.ResponseWriter, r *http.Request, log *l.Logger, tmp *t.Template, coll *mongo.Collection) {

	// handle method
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	// load workouts from db
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

	var workouts []models.Workout
	if err = cursor.All(ctx, &workouts); err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

	// analyze workouts
	data := models.Analyze(workouts, log)

	// return template
	if err := tmp.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

}
