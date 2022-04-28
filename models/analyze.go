package models

import (
	"time"
)

func Analyize(workouts []Workout) (Analytics, error) {
	var data Analytics

	durTotal, err := time.ParseDuration("260h27m")
	if err != nil {
		return data, err
	}

	durAvg, err := time.ParseDuration("38m")
	if err != nil {
		return data, err
	}

	durPeak, err := time.ParseDuration("10h17m")
	if err != nil {
		return data, err
	}

	data = Analytics{
		Total: Total{
			Workouts:  119,
			Distance:  340.0,
			Elevation: 1730,
			Duration:  durTotal,
		},
		Average: Average{
			Duration: durAvg,
			Pace:     6.37,
			HR:       137,
		},
		Peak: Peak{
			Duration: durPeak,
			Pace:     4.14,
			HR:       192,
		},
		Terrain: TerrainAnalytics{
			Road:      52.0,
			Trail:     31.0,
			Treadmill: 17.0,
		},
		Footwear: FootwearAnalytics{
			Barefoot: 3.0,
			Minimal:  12.0,
			Standard: 85,
		},
	}

	return data, nil
}
