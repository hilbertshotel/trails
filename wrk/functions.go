package wrk

import (
	"fmt"
	"time"
	"trails/dep"
)

// CALCULATE TOTALS
func (workouts Workouts) CalcTotals(log *dep.Logger) (Total, error) {
	numOfWorkouts := len(workouts)

	// Handle no workouts scenario
	if numOfWorkouts == 0 {
		return Total{
			Days:      0,
			Workouts:  0,
			Distance:  "0",
			Duration:  "0",
			Elevation: 0,
		}, nil
	}

	// Else
	distance, elevation, minutes, dates := 0.0, 0, 0.0, []string{}

	for _, workout := range workouts {
		distance += workout.Distance
		elevation += workout.Elevation
		dates = append(dates, workout.Date)

		duration, err := time.ParseDuration(workout.Duration)
		if err != nil {
			log.Error(err)
			return Total{}, err
		}
		minutes += duration.Minutes()
	}

	return Total{
		Days:      parseDays(dates[0]),
		Workouts:  numOfWorkouts,
		Streak:    parseStreak(dates),
		Distance:  formatDistance(distance),
		Duration:  formatDuration(minutes),
		Elevation: elevation,
	}, nil
}

// PARSE DAYS
func parseDays(date string) int {
	start, _ := time.Parse(time.RFC822, date+" 10:00 MST")
	end := time.Now()
	return int(end.Sub(start).Hours() / 24)
}

// FORMAT DURATION
func formatDuration(minutes float64) string {
	mins := int(minutes)
	if mins < 60 {
		return fmt.Sprintf("%vm", mins)
	} else if mins < 1440 {
		h := mins / 60
		m := mins % 60
		if m == 0 {
			return fmt.Sprintf("%vh", h)
		}
		return fmt.Sprintf("%vh%vm", h, m)
	} else {
		d := mins / 1440
		rd := mins % 1440
		h := rd / 60
		m := rd % 60
		if m == 0 && h == 0 {
			return fmt.Sprintf("%vd", d)
		} else if m == 0 {
			return fmt.Sprintf("%vd%vh", d, h)
		} else if h == 0 {
			return fmt.Sprintf("%vd%vm", d, m)
		}
		return fmt.Sprintf("%vd%vh%vm", d, h, m)
	}
}

// PARSE STREAK
func parseStreak(dates []string) int {

	if len(dates) == 1 {
		return 1
	}

	longest, streak := 0, 1
	prev, _ := time.Parse(time.RFC822, dates[0]+" 10:00 MST")

	for i := 1; i < len(dates); i++ {
		current, _ := time.Parse(time.RFC822, dates[i]+" 10:00 MST")

		if prev.Add(24*time.Hour).Equal(current) || prev.Equal(current) {
			streak++
		} else {
			streak = 1
		}

		if longest < streak {
			longest = streak
		}

		prev = current
	}

	return longest
}

// FORMAT DISTANCE
func formatDistance(distance float64) string {
	return fmt.Sprintf("%1.f", distance)
}
