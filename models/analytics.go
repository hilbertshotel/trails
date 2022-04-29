package models

type Peak struct {
	Duration  string
	Pace      string
	Distance  float64
	Elevation int
}

type Total struct {
	Workouts  int
	Distance  float64
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
