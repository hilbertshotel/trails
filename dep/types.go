package dep

import (
	"database/sql"
	"html/template"
)

// LOGGER
type Logger struct {
	UserError func(string)
	Error     func(error)
	Ok        func(string)
}

// CONFIG
type Config struct {
	HostAddr string
	Template string
	ConnStr  string
}

// DEPENDENCIES
type Dependencies struct {
	Log *Logger
	Cfg *Config
	DB  *sql.DB
	Tmp *template.Template
}
