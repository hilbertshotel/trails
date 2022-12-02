package wrk

import (
	"sync"
	"time"
	"trails/lib"
	"trails/logger"
)

// ANALYZE WORKOUTS
func (workouts Workouts) Analyze(log *logger.Logger) Analytics {

	// HANDLE NO WORKOUTS
	if len(workouts) == 0 {
		return Analytics{
			Total: Total{
				Workouts:  0,
				Distance:  "0",
				Elevation: 0,
				Duration:  "0m",
				Range:     0,
				Streak:    0,
			},
			Peak: Peak{
				Duration:  "0m",
				Pace:      "0",
				Distance:  0,
				Elevation: 0,
			},
			Terrain: TerrainAnalytics{
				Road:  "0%",
				Trail: "0%",
				Gym:   "0%",
			},
			Footwear: FootwearAnalytics{
				Minimal:  "0%",
				Standard: "0%",
			},
		}
	}

	// HANDLE MULTIPLE WORKOUTS
	var wg sync.WaitGroup

	// calculate totals
	var total Total

	wg.Add(1)
	go func(total *Total, wg *sync.WaitGroup) {
		defer wg.Done()

		var minutes float64
		var dates []string
		var dist float64

		for _, w := range workouts {
			total.Workouts++
			dist += w.Distance
			total.Elevation += w.Elev.Gain

			duration, err := time.ParseDuration(w.Duration)
			if err != nil {
				log.Error(err)
				return
			}

			minutes += duration.Minutes()
			dates = append(dates, w.Date)
		}

		total.Distance = lib.FormatDistance(dist)
		total.Duration = lib.FormatDuration(minutes)
		total.Range = lib.ParseRange(dates[0])
		total.Streak = lib.ParseStreak(dates)

	}(&total, &wg)

	// calculate peaks
	var peak Peak

	wg.Add(1)
	go func(peak *Peak, wg *sync.WaitGroup) {
		defer wg.Done()

		pace := workouts[0].Pace.Best
		dist, dur, elev := 0.0, 0.0, 0
		for _, w := range workouts {
			duration, err := time.ParseDuration(w.Duration)
			if err != nil {
				log.Error(err)
				return
			}
			if duration.Minutes() > dur {
				dur = duration.Minutes()
			}
			if w.Pace.Best < pace {
				pace = w.Pace.Best
			}
			if w.Distance > dist {
				dist = w.Distance
			}
			if w.Elev.Gain > elev {
				elev = w.Elev.Gain
			}
		}

		peak.Duration = lib.FormatDuration(dur)
		peak.Pace = lib.FormatPace(pace)
		peak.Distance = dist
		peak.Elevation = elev

	}(&peak, &wg)

	// calculate terrain
	var terrain TerrainAnalytics

	wg.Add(1)
	go func(terrain *TerrainAnalytics, wg *sync.WaitGroup) {
		defer wg.Done()

		road, trail, gym := 0.0, 0.0, 0.0
		for _, w := range workouts {
			switch w.Location.Terrain {
			case Road:
				road++
			case Trail:
				trail++
			case Gym:
				gym++
			}
		}

		total := float64(road + trail + gym)
		terrain.Road = lib.FormatPercent(road / total * 100)
		terrain.Trail = lib.FormatPercent(trail / total * 100)
		terrain.Gym = lib.FormatPercent(gym / total * 100)

	}(&terrain, &wg)

	// calculate footwear
	var footwear FootwearAnalytics

	wg.Add(1)
	go func(footwear *FootwearAnalytics, wg *sync.WaitGroup) {
		defer wg.Done()

		mnml, stnd := 0.0, 0.0
		for _, w := range workouts {
			switch w.Footwear {
			case Minimal:
				mnml++
			case Standard:
				stnd++
			}
		}

		total := float64(mnml + stnd)
		footwear.Minimal = lib.FormatPercent(mnml / total * 100)
		footwear.Standard = lib.FormatPercent(stnd / total * 100)

	}(&footwear, &wg)

	wg.Wait()

	return Analytics{
		Total:    total,
		Peak:     peak,
		Terrain:  terrain,
		Footwear: footwear,
	}
}
