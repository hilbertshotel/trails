package handlers

import (
	"net/http"
	"text/template"
	"time"
	"trails/logger"
	"trails/models"
)

func get(w http.ResponseWriter, r *http.Request, log *logger.Logger, tmp *template.Template) {

	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	// load workouts from database
	// apply business logic to condense data
	durTotal, err := time.ParseDuration("260h27m")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

	durAvg, err := time.ParseDuration("38m")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

	durPeak, err := time.ParseDuration("10h17m")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

	data := models.Analytics{
		Total: models.Total{
			Workouts:  119,
			Distance:  340.0,
			Elevation: 1730,
			Duration:  durTotal,
		},
		Average: models.Average{
			Duration: durAvg,
			Pace:     6.37,
			HR:       137,
		},
		Peak: models.Peak{
			Duration: durPeak,
			Pace:     4.14,
			HR:       192,
		},
		Terrain: models.TerrainAnalytics{
			Road:      52.0,
			Trail:     31.0,
			Treadmill: 17.0,
		},
		Footwear: models.FootwearAnalytics{
			Barefoot: 3.0,
			Minimal:  12.0,
			Standard: 85,
		},
	}

	if err := tmp.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

}
