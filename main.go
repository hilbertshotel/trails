package main

import (
	"database/sql"
	"net/http"
	"text/template"
	"time"
	"trails/config"
	"trails/handlers"
	"trails/logger"

	_ "github.com/lib/pq"
)

const ADDRESS = "127.0.0.1:9990"

func main() {

	// logging
	log := logger.Init()

	// config
	cfg := config.Init()

	// database
	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		log.Error(err)
		return
	}

	// templates
	tmp, err := template.ParseGlob(cfg.Template)
	if err != nil {
		log.Error(err)
		return
	}

	// server
	server := http.Server{
		Addr:         cfg.HostAddr,
		Handler:      handlers.Mux(log, tmp, db),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// channels
	ech := make(chan error)

	// start
	go func(ch chan error) {
		log.Ok("Listening on " + cfg.HostAddr)
		ch <- server.ListenAndServe()
	}(ech)

	// shutdown
	log.Error(<-ech)
}
