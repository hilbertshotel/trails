package wrk

import (
	"database/sql"
	"sort"
)

func Load(db *sql.DB) (Workouts, error) {
	// load workouts from database
	query := "SELECT * FROM workouts"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workouts Workouts
	for rows.Next() {
		var workout Workout

		err = rows.Scan(&workout.Id, &workout.Date, &workout.Distance, &workout.Duration.Back,
			&workout.Elevation, &workout.AvgPace, &workout.AvgHR)
		if err != nil {
			return nil, err
		}
		workouts = append(workouts, workout)
	}

	// sort workouts by id
	sort.Slice(workouts, func(i, j int) bool {
		return workouts[i].Id < workouts[j].Id
	})

	return workouts, nil
}
