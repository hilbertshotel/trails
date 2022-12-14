package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

func adder(filename string) {
	// open file
	file, err := os.Open(filename)
	handleError(err)
	defer file.Close()

	// read file
	raw, err := ioutil.ReadAll(file)
	handleError(err)

	// unmarshal file
	var workouts Workouts
	err = json.Unmarshal(raw, &workouts)
	handleError(err)

	// open db
	db, err := sql.Open("postgres", CONN_STR)
	handleError(err)
	defer db.Close()

	// insert in db
	for _, workout := range workouts {
		_, err = db.Exec(`INSERT INTO workouts
		(date, distance, duration, elevation, avg_pace, avg_hr)
		VALUES ($1, $2, $3, $4, $5, $6)`, workout.Date, workout.Distance,
			workout.Duration, workout.Elevation, workout.AvgPace, workout.AvgHR)
		handleError(err)
	}

	fmt.Println("ok")
}
