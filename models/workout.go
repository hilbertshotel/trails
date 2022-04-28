package models

import "time"

type Pace struct {
	Avg  float64
	Best float64
}

type HR struct {
	Avg int
	Max int
}

type Elev struct {
	Gain int
	Loss int
}

type Terrain string

const (
	Road      Terrain = "road"
	Trail     Terrain = "trail"
	Treadmill Terrain = "treadmill"
)

type Location struct {
	Name    string
	Terrain Terrain
}

type Footwear string

const (
	Barefoot Footwear = "barefoot"
	Minimal  Footwear = "minimal"
	Standard Footwear = "standard"
)

type Workout struct {
	Duration time.Duration
	Distance float64
	Pace     Pace
	HR       HR
	Elev     Elev
	Terrain  Terrain
	Footwear Footwear
}
