package models

import "time"

type Peak struct {
	Duration time.Duration
	Pace     float64
	HR       int
}

type Average struct {
	Duration time.Duration
	Pace     float64
	HR       int
}

type Total struct {
	Workouts  int
	Distance  float64
	Elevation int
	Duration  time.Duration
}

type TerrainAnalytics struct {
	Road      float64
	Trail     float64
	Treadmill float64
}

type FootwearAnalytics struct {
	Barefoot float64
	Minimal  float64
	Standard float64
}

type Analytics struct {
	Total    Total
	Average  Average
	Peak     Peak
	Terrain  TerrainAnalytics
	Footwear FootwearAnalytics
}
