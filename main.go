package main

import (
	"net/http"
	"os"
	"time"
	"trails/dep"
	"trails/handlers"
)

func main() {

	// get config type
	args := os.Args[1:]
	if len(args) != 1 {
		panic("No config path argument provided!")
	}

	// Load Dependencies
	d, err := dep.Load(args[0])
	if err != nil {
		panic(err)
	}
	defer d.DB.Close()

	// Create Server
	server := http.Server{
		Addr:         d.Cfg.HostAddr,
		Handler:      handlers.Mux(d),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Channels
	ech := make(chan error)

	// Start Server
	go func(ch chan error) {
		d.Log.Ok("Service listening @ " + d.Cfg.HostAddr)
		ch <- server.ListenAndServe()
	}(ech)

	// Shutdown
	d.Log.Error(<-ech)
}
