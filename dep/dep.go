package dep

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// LOAD DEPENDENCIES
func Load() (*Dependencies, error) {
	log := initLogger()
	cfg := initConfig()
	coll, client, err := initDatabase(cfg)
	if err != nil {
		return nil, err
	}
	tmp, err := initTemplates(cfg)
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		Log:    log,
		Cfg:    cfg,
		Client: client,
		Coll:   coll,
		Tmp:    tmp,
	}, nil
}

// INITIALIZE LOGGER
func initLogger() *Logger {
	okLog := log.New(os.Stdout, "OK ", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime)

	return &Logger{
		Error: func(err error) {
			_, file, line, _ := runtime.Caller(1)
			msg := fmt.Sprintf("(%v %v) %v", filepath.Base(file), line, err)
			errLog.Println(msg)
		},

		Ok: func(msg string) {
			okLog.Println(msg)
		},
	}
}

// INITIALIZE DATABASE
func initDatabase(cfg *Config) (*mongo.Collection, *mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo.Uri))
	if err != nil {
		return nil, nil, err
	}

	coll := client.Database(cfg.Mongo.Database).Collection(cfg.Mongo.Collection)
	return coll, client, nil
}

// INITIALIZE TEMPLATES
func initTemplates(cfg *Config) (*template.Template, error) {
	tmp, err := template.ParseGlob(cfg.Template)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}
