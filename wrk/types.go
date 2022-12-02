package wrk

// WORKOUT
type Pace struct {
	Avg  float64 `bson:"avg"`
	Best float64 `bson:"best"`
}

type HR struct {
	Avg int `bson:"avg"`
	Max int `bson:"max"`
}

type Elev struct {
	Gain int `bson:"gain"`
	Loss int `bson:"loss"`
}

type Terrain string

const (
	Road  Terrain = "road"
	Trail Terrain = "trail"
	Gym   Terrain = "gym"
)

type Location struct {
	Name    string  `bson:"name"`
	Terrain Terrain `bson:"terrain"`
}

type Footwear string

const (
	Minimal  Footwear = "minimal"
	Standard Footwear = "standard"
)

type Workout struct {
	Date     string   `bson:"date"`
	Duration string   `bson:"duration"`
	Distance float64  `bson:"distance"`
	Pace     Pace     `bson:"pace"`
	HR       HR       `bson:"hr"`
	Elev     Elev     `bson:"elev"`
	Location Location `bson:"location"`
	Footwear Footwear `bson:"footwear"`
}

type Workouts []Workout

// ANALYTICS
type Peak struct {
	Duration  string
	Pace      string
	Distance  float64
	Elevation int
}

type Total struct {
	Workouts  int
	Distance  string
	Elevation int
	Duration  string
	Range     int
	Streak    int
}

type TerrainAnalytics struct {
	Road  string
	Trail string
	Gym   string
}

type FootwearAnalytics struct {
	Minimal  string
	Standard string
}

type Analytics struct {
	Total    Total
	Peak     Peak
	Terrain  TerrainAnalytics
	Footwear FootwearAnalytics
}
