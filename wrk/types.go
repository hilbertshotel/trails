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

// WORKOUT
type Workout struct {
	Date      string  `bson:"date"`
	Distance  float64 `bson:"distance"`
	Duration  int     `bson:"duration"`
	Elevation int     `bson:"elevation"`
	AvgPace   float64 `bson:"avg_pace"`
	AvgHR     int     `bson:"avg_hr"`
}

type Workouts []Workout
