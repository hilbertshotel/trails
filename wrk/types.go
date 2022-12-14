package wrk

// TOTALS
type Total struct {
	Days      int
	Workouts  int
	Streak    int
	Distance  string
	Duration  string
	Elevation int
}

// WORKOUT BACKEND
type Duration struct {
	Front float64 `json:"front"`
	Back  string  `json:"back"`
}

type Workout struct {
	Id        int      `json:"id"`
	Date      string   `json:"date"`
	Distance  float64  `json:"distance"`
	Duration  Duration `json:"duration"`
	Elevation int      `json:"elevation"`
	AvgPace   float64  `json:"avgPace"`
	AvgHR     int      `json:"avgHR"`
}

type Workouts []Workout
