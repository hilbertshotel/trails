package wrk

func (workouts Workouts) Sort(args map[string]string) Workouts {

	// make copy of workouts slice
	sortedWorkouts := make([]Workout, len(workouts))
	copy(sortedWorkouts, workouts)

	switch args["name"] {

	// SORT BY DATE
	case "date":

		switch args["direction"] {
		case "asc":
			return workouts
		case "desc":
			return workouts.reverse()
		}

	// SORT BY NUMBER
	case "distance", "elevation", "pace", "hr":

	// SORT BY DURATION
	case "duration":
	}

	return nil
}

// REVERSE WORKOUTS
func (workouts Workouts) reverse() Workouts {
	wrkLen := len(workouts)
	reversedWorkouts := make(Workouts, wrkLen)

	for i, workout := range workouts {
		j := wrkLen - i - 1
		reversedWorkouts[j] = workout
	}

	return reversedWorkouts
}
