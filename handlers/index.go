package handlers

import (
	"net/http"
	"trails/dep"
	"trails/wrk"
)

func index(w http.ResponseWriter, r *http.Request, d *dep.Dependencies) {

	// handle method
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	// load workouts from db
	workouts, err := wrk.Load(d.DB)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}

	// get totals
	totals, err := workouts.CalcTotals(d.Log)
	if err != nil {
		d.Log.Error(err)
		return
	}

	// return template
	if err := d.Tmp.ExecuteTemplate(w, "index.html", totals); err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}

}
