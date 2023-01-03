package main

import (
	"fmt"
	"os"
)

type Workout struct {
	Id        int     `json:"id"`
	Date      string  `json:"date"`
	Distance  float64 `json:"distance"`
	Duration  string  `json:"duration"`
	Elevation int     `json:"elevation"`
	AvgPace   float64 `json:"avg_pace"`
	AvgHR     int     `json:"avg_hr"`
}
type Workouts []Workout

const CONN_STR = "user=postgres dbname=trails password=postgres sslmode=disable"

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// handle arguments
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Wrong number of arguments")
		os.Exit(1)
	}

	// handle commands
	switch args[0] {
	case "add":
		adder(args[1])
	case "dump":
		dumper(args[1])
	default:
		fmt.Println("Unknown command:", args[0])
	}
}
