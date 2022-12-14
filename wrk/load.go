package wrk

import "database/sql"

func Load(db *sql.DB) (Workouts, error) {
	query := "SELECT date, distance, duration, elevation, avg_pace, avg_hr FROM workouts"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workouts Workouts
	for rows.Next() {
		var workout Workout

		err = rows.Scan(&workout.Date, &workout.Distance, &workout.Duration,
			&workout.Elevation, &workout.AvgPace, &workout.AvgHR)
		if err != nil {
			return nil, err
		}
		workouts = append(workouts, workout)
	}

	return workouts, nil
}
