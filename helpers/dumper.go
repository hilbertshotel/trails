package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func dumper(filename string) {
	// open db
	db, err := sql.Open("postgres", CONN_STR)
	handleError(err)
	defer db.Close()

	// load workouts from database
	query := "SELECT * FROM workouts"
	rows, err := db.Query(query)
	handleError(err)
	defer rows.Close()

	var workouts Workouts
	for rows.Next() {
		var workout Workout

		err = rows.Scan(&workout.Id, &workout.Date, &workout.Distance, &workout.Duration,
			&workout.Elevation, &workout.AvgPace, &workout.AvgHR)
		handleError(err)

		workouts = append(workouts, workout)
	}

	// sort workouts by id
	sort.Slice(workouts, func(i, j int) bool {
		return workouts[i].Id < workouts[j].Id
	})

	// marshal workouts
	jsonData, err := json.Marshal(workouts)
	handleError(err)

	// open file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0600)
	handleError(err)
	defer file.Close()

	// write to file
	_, err = file.Write(jsonData)
	handleError(err)

	fmt.Println("ok")
}
