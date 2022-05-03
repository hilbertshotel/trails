package model

// BASICS
type Pace struct {
	Avg  float64 `json:"avg"`
	Best float64 `json:"best"`
}

type HR struct {
	Avg int `json:"avg"`
	Max int `json:"max"`
}

type Elev struct {
	Gain int `json:"gain"`
	Loss int `json:"loss"`
}

// LOCATION
type Terrain string

const (
	Road  Terrain = "road"
	Trail Terrain = "trail"
	Gym   Terrain = "gym"
)

type Location struct {
	Name    string  `json:"name"`
	Terrain Terrain `json:"terrain"`
}

// FOOTWEAR
type Footwear string

const (
	Minimal  Footwear = "minimal"
	Standard Footwear = "standard"
)

// WORKOUT
type Workout struct {
	Date     string   `json:"date"`
	Duration string   `json:"duration"`
	Distance float64  `json:"distance"`
	Pace     Pace     `json:"pace"`
	HR       HR       `json:"hr"`
	Elev     Elev     `json:"elev"`
	Location Location `json:"location"`
	Footwear Footwear `json:"footwear"`
}
