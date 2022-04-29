package models

// BASICS
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

// LOCATION
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

// FOOTWEAR
type Footwear string

const (
	Minimal  Footwear = "minimal"
	Standard Footwear = "standard"
)

// WORKOUT
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
