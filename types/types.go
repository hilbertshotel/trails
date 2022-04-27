package types

import "time"

// Pace
type Pace struct {
	Avg  float64
	Best float64
}

// Heart Rate
type HR struct {
	Avg float64
	Max float64
}

// Elevation
type Elev struct {
	Gain int
	Loss int
}

// Terrain Type
type Terrain int

const (
	Road Terrain = iota
	Trail
	Treadmill
)

// Session
type Session struct {
	Id       int
	Date     time.Time
	Duration time.Duration
	Distance float64
	Pace     Pace
	HR       HR
	Elev     Elev
	Terrain  Terrain
}
