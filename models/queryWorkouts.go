package models

import (
	"database/sql"
)

func QueryWorkouts(db *sql.DB) ([]Workout, error) {
	var workouts []Workout

	rows, err := db.Query(`SELECT duration, distance, pace_avg, pace_best,
	hr_avg, hr_max, elev_gain, elev_loss, terrain, footwear FROM workouts;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var wk Workout
		err = rows.Scan(&wk.Duration, &wk.Distance, &wk.Pace.Avg, &wk.Pace.Best,
			&wk.HR.Avg, &wk.HR.Max, &wk.Elev.Gain, &wk.Elev.Loss, &wk.Terrain, &wk.Footwear)
		if err != nil {
			return nil, err
		}
		workouts = append(workouts, wk)
	}

	return workouts, err
}
