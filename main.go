package main

import (
	"net/http"
	"time"
	"trails/dep"
	"trails/handlers"
)

func main() {

	// Load Dependencies
	d, err := dep.Load()
	if err != nil {
		panic(err)
	}
	defer d.DB.Close()

	// server
	server := http.Server{
		Addr:         d.Cfg.HostAddr,
		Handler:      handlers.Mux(d),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// channels
	ech := make(chan error)

	// start
	go func(ch chan error) {
		d.Log.Ok("Service listening @ " + d.Cfg.HostAddr)
		ch <- server.ListenAndServe()
	}(ech)

	// shutdown
	d.Log.Error(<-ech)
}
