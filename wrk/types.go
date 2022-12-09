package wrk

// TOTALS
type Total struct {
	Days      int
	Workouts  int
	Streak    int
	Distance  int
	Duration  string
	Elevation int
}

// WORKOUT
type Workout struct {
	Date      string  `bson:"date"`
	Distance  int     `bson:"distance"`
	Duration  string  `bson:"duration"`
	Elevation int     `bson:"elevation"`
	AvgPace   float64 `bson:"avg_pace"`
	AvgHR     int     `bson:"avg_hr"`
}

type Workouts []Workout

// DATA
type Data struct {
	Total    Total
	Best     Workouts
	Workouts Workouts
}
