package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

type Workout struct {
	Date      string  `json:"date"`
	Distance  float64 `json:"distance"`
	Duration  string  `json:"duration"`
	Elevation int     `json:"elevation"`
	AvgPace   float64 `json:"avg_pace"`
	AvgHR     int     `json:"avg_hr"`
}
type Workouts []Workout

const (
	JSON_FILE = "workouts.json"
	CONN_STR  = "user=postgres dbname=trails password=postgres sslmode=disable"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	// open file
	file, err := os.Open(JSON_FILE)
	handleError(err)
	defer file.Close()

	// read file
	raw, err := ioutil.ReadAll(file)
	handleError(err)

	// unmarshal file
	var jsonData map[string]Workouts
	err = json.Unmarshal(raw, &jsonData)
	handleError(err)
	workouts := jsonData["workouts"]

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
