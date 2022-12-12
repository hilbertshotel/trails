package handlers

import (
	"context"
	"net/http"
	"time"
	"trails/dep"
	"trails/wrk"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
)

func sorter(w http.ResponseWriter, r *http.Request, d *dep.Dependencies) {

	// UPGRADE HTTP HANDLER TO WS SERVER
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}
	defer conn.Close()

	// START SERVER
	for {

		// Wait for sorting argument
		_, arg, err := conn.ReadMessage()
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			d.Log.Error(err)
			return
		}

		// Load workouts from db
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

		// Sort workouts by argument
		asc, desc := workouts.Sort(string(arg))
		if asc == nil {
			http.Error(w, "Internal Server Error", 500)
			d.Log.UserError("Unsupported sorting argument")
			return
		}

		data := map[string]wrk.Workouts{
			"asc":  asc,
			"desc": desc,
		}

		// Return sorted data as JSON
		err = conn.WriteJSON(data)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			d.Log.Error(err)
			return
		}
	}
}
