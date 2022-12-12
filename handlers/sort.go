package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"trails/dep"
	"trails/wrk"

	"go.mongodb.org/mongo-driver/bson"
)

func sort(w http.ResponseWriter, r *http.Request, d *dep.Dependencies) {

	// handle method
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	// get sorting argument
	args := strings.Split(r.URL.Path[1:], "/")[1:]
	if len(args) != 1 {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	arg := args[0]

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

	// sort workouts by argument
	asc, desc := workouts.Sort(arg)
	if asc == nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.UserError("Unsupported sorting argument")
		return
	}

	data := map[string]wrk.Workouts{
		"asc":  asc,
		"desc": desc,
	}

	// marshal workouts
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}

	// respond
	w.Header().Set("content-type", "application/json")
	w.Write(response)
}
