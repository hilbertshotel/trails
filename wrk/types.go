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
	Duration  string  `bson:"duration"`
	Elevation int     `bson:"elevation"`
	AvgPace   float64 `bson:"avg_pace"`
	AvgHR     int     `bson:"avg_hr"`
}

type Workouts []Workout

// SORTED DATA
type SortedData struct {
	DateAsc       Workouts
	DateDesc      Workouts
	DistanceAsc   Workouts
	DistanceDesc  Workouts
	DurationAsc   Workouts
	DurationDesc  Workouts
	ElevationAsc  Workouts
	ElevationDesc Workouts
	PaceAsc       Workouts
	PaceDesc      Workouts
	HRAsc         Workouts
	HRDesc        Workouts
}
