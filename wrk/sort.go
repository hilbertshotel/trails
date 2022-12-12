package wrk

import (
	"sort"
	"time"
)

func (workouts Workouts) Sort(argument string) (Workouts, Workouts) {

	// make slice copies
	asc := make([]Workout, len(workouts))
	copy(asc, workouts)

	desc := make([]Workout, len(workouts))
	copy(desc, workouts)

	switch argument {

	// SORT BY DATE
	case "date":
		sort.Slice(asc, func(i, j int) bool {
			a, _ := time.Parse("02 Jan 06", workouts[i].Date)
			b, _ := time.Parse("02 Jan 06", workouts[j].Date)
			return a.Before(b)
		})
		sort.Slice(desc, func(i, j int) bool {
			a, _ := time.Parse("02 Jan 06", workouts[j].Date)
			b, _ := time.Parse("02 Jan 06", workouts[i].Date)
			return a.Before(b)
		})
		return asc, desc

	// SORT BY NUMBER
	case "distance", "elevation", "pace", "hr":

	// SORT BY DURATION
	case "duration":
	}

	return nil, nil
}
