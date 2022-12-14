package wrk

import "time"

func (workouts *Workouts) ParseDuration() error {
	for i, workout := range *workouts {
		duration, err := time.ParseDuration(workout.Duration.Back)
		if err != nil {
			return err
		}
		(*workouts)[i].Duration.Front = duration.Seconds()
	}
	return nil
}
