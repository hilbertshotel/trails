package dep

import (
	"html/template"

	"go.mongodb.org/mongo-driver/mongo"
)

// LOGGER
type Logger struct {
	Error func(error)
	Ok    func(string)
}

// CONFIG
type Mongo struct {
	Uri        string
	Database   string
	Collection string
}

type Config struct {
	HostAddr string
	Template string
	Mongo    Mongo
}

// DEPENDENCIES
type Dependencies struct {
	Log  *Logger
	Cfg  *Config
	Coll *mongo.Collection
	Tmp  *template.Template
}
