package models

import "time"

// Pace
type Pace struct {
	Avg  float64
	Best float64
}

// Heart Rate
type HR struct {
	Avg int
	Max int
}

// Elevation
type Elev struct {
	Gain int
	Loss int
}

// Terrain Type
type Terrain string

const (
	Road      Terrain = "road"
	Trail     Terrain = "trail"
	Treadmill Terrain = "treadmill"
)

// Location
type Location struct {
	Name    string
	Terrain Terrain
}

// Footwear
type Footwear string

const (
	Barefoot Footwear = "barefoot"
	Minimal  Footwear = "minimal"
	Standard Footwear = "standard"
)

// Session
type Workout struct {
	Id       int
	Date     time.Time
	Duration time.Duration
	Distance float64
	Pace     Pace
	HR       HR
	Elev     Elev
	Location Location
	Footwear Footwear
}
