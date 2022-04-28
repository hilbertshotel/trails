package main

import (
	"context"
	"net/http"
	"text/template"
	"time"
	"trails/config"
	"trails/handlers"
	"trails/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// logging
	log := logger.Init()

	// config
	cfg := config.Init()

	// database
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo.Uri))
	if err != nil {
		log.Error(err)
		return
	}
	defer client.Disconnect(ctx)

	coll := client.Database(cfg.Mongo.Database).Collection(cfg.Mongo.Coll)

	// templates
	tmp, err := template.ParseGlob(cfg.Template)
	if err != nil {
		log.Error(err)
		return
	}

	// server
	server := http.Server{
		Addr:         cfg.HostAddr,
		Handler:      handlers.Mux(log, tmp, coll),
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
