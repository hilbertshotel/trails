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
	HostAddr string `json:"hostAddr"`
	TmpDir   string `json:"tmpDir"`
	ConnStr  string `json:"connStr"`
}

// DEPENDENCIES
type Dependencies struct {
	Log *Logger
	Cfg *Config
	DB  *sql.DB
	Tmp *template.Template
}
