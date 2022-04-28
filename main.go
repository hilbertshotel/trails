package main

import (
	"net/http"
	"text/template"
	"time"
	"trails/config"
	"trails/handlers"
	"trails/logger"
)

const ADDRESS = "127.0.0.1:9990"

func main() {

	// logging
	log := logger.Init()

	// config
	cfg := config.Init()

	// templates
	tmp, err := template.ParseGlob(cfg.Template)
	if err != nil {
		log.Error(err)
	}

	// server
	server := http.Server{
		Addr:         cfg.HostAddr,
		Handler:      handlers.Mux(log, tmp),
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
