package models

import (
	"fmt"
	"sync"
	"time"
	"trails/logger"
)

func Analyze(workouts []Workout, log *logger.Logger) Analytics {
	var data Analytics

	// HANDLE NO WORKOUTS
	if len(workouts) == 0 {

		data = Analytics{
			Total: Total{
				Workouts:  0,
				Distance:  0,
				Elevation: 0,
				Duration:  "0",
			},
			Average: Average{
				Duration: "0",
				Pace:     0,
				HR:       0,
			},
			Peak: Peak{
				Duration: "0",
				Pace:     0,
				HR:       0,
			},
			Terrain: TerrainAnalytics{
				Road:      0,
				Trail:     0,
				Treadmill: 0,
			},
			Footwear: FootwearAnalytics{
				Barefoot: 0,
				Minimal:  0,
				Standard: 0,
			},
		}

		return data
	}

	// HANDLE SINGLE WORKOUT
	if len(workouts) == 1 {
		workout := workouts[0]

		data = Analytics{
			Total: Total{
				Workouts:  1,
				Distance:  workout.Distance,
				Elevation: workout.Elev.Gain,
				Duration:  workout.Duration,
			},
			Average: Average{
				Duration: workout.Duration,
				Pace:     workout.Pace.Avg,
				HR:       workout.HR.Avg,
			},
			Peak: Peak{
				Duration: workout.Duration,
				Pace:     workout.Pace.Best,
				HR:       workout.HR.Max,
			},
		}

		switch workout.Location.Terrain {
		case Road:
			data.Terrain.Road = 100
		case Trail:
			data.Terrain.Trail = 100
		case Treadmill:
			data.Terrain.Treadmill = 100
		}

		switch workout.Footwear {
		case Barefoot:
			data.Footwear.Barefoot = 100
		case Minimal:
			data.Footwear.Minimal = 100
		case Standard:
			data.Footwear.Standard = 100
		}

		return data
	}

	// HANDLE MULTIPLE WORKOUTS
	var wg sync.WaitGroup

	// calculate totals
	var total Total

	wg.Add(1)
	go func(total *Total, wg *sync.WaitGroup) {
		defer wg.Done()

		var minutes float64
		for _, w := range workouts {
			total.Workouts++
			total.Distance += w.Distance
			total.Elevation += w.Elev.Gain

			dur, err := time.ParseDuration(w.Duration)
			if err != nil {
				log.Error(err)
				return
			}
			minutes += dur.Minutes()
		}

		// convert minutes to string 133d16h1m
		total.Duration = fmt.Sprintf("%vm", minutes) //////////////////////

	}(&total, &wg)

	// calculate averages
	var average Average

	wg.Add(1)
	go func(average *Average, wg *sync.WaitGroup) {
		defer wg.Done()

		minutes, pace, hr := 0.0, 0.0, 0
		for _, w := range workouts {
			pace += w.Pace.Avg
			hr += w.HR.Avg

			dur, err := time.ParseDuration(w.Duration)
			if err != nil {
				log.Error(err)
				return
			}
			minutes += dur.Minutes()
		}

		length := float64(len(workouts))
		average.Duration = fmt.Sprintf("%vm", minutes/length) ////////////////////
		average.Pace = pace / length
		average.HR = int(float64(hr) / length)

	}(&average, &wg)

	// calculate peaks
	var peak Peak

	wg.Add(1)
	go func(peak *Peak, wg *sync.WaitGroup) {
		defer wg.Done()

		maxDur, bestPace, maxHR := 0.0, 0.0, 0
		for _, w := range workouts {
			dur, err := time.ParseDuration(w.Duration)
			if err != nil {
				log.Error(err)
				return
			}
			if dur.Minutes() > maxDur {
				maxDur = dur.Minutes()
			}
			if w.Pace.Best > bestPace {
				bestPace = w.Pace.Best
			}
			if w.HR.Max > maxHR {
				maxHR = w.HR.Max
			}
		}

		peak.Duration = fmt.Sprintf("%vm", maxDur) ////////////////
		peak.Pace = bestPace
		peak.HR = maxHR

	}(&peak, &wg)

	// calculate terrain
	var terrain TerrainAnalytics

	wg.Add(1)
	go func(terrain *TerrainAnalytics, wg *sync.WaitGroup) {
		defer wg.Done()

		roadCount, trailCount, treadCount := 0.0, 0.0, 0.0
		for _, w := range workouts {
			switch w.Location.Terrain {
			case Road:
				roadCount++
			case Trail:
				trailCount++
			case Treadmill:
				treadCount++
			}
		}

		total := roadCount + trailCount + treadCount
		terrain.Road = roadCount / total
		terrain.Trail = trailCount / total
		terrain.Treadmill = treadCount / total

	}(&terrain, &wg)

	// calculate footwear
	var footwear FootwearAnalytics

	wg.Add(1)
	go func(footwear *FootwearAnalytics, wg *sync.WaitGroup) {
		defer wg.Done()

		bareCount, minCount, standCount := 0.0, 0.0, 0.0
		for _, w := range workouts {
			switch w.Footwear {
			case Barefoot:
				bareCount++
			case Minimal:
				minCount++
			case Standard:
				standCount++
			}
		}

		total := bareCount + minCount + standCount
		footwear.Barefoot = bareCount / total
		footwear.Minimal = minCount / total
		footwear.Standard = standCount / total

	}(&footwear, &wg)

	wg.Wait()

	data.Total = total
	data.Average = average
	data.Peak = peak
	data.Terrain = terrain
	data.Footwear = footwear

	return data
}
