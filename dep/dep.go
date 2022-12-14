package dep

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"runtime"

	_ "github.com/lib/pq"
)

// LOAD DEPENDENCIES
func Load() (*Dependencies, error) {
	log := initLogger()
	cfg := initConfig()
	db, err := initDatabase(cfg)
	if err != nil {
		return nil, err
	}
	tmp, err := initTemplates(cfg)
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		Log: log,
		Cfg: cfg,
		DB:  db,
		Tmp: tmp,
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

		UserError: func(err string) {
			_, file, line, _ := runtime.Caller(1)
			msg := fmt.Sprintf("(%v %v) %v", filepath.Base(file), line, err)
			errLog.Println(msg)
		},
	}
}

// INITIALIZE DATABASE
func initDatabase(cfg *Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// INITIALIZE TEMPLATES
func initTemplates(cfg *Config) (*template.Template, error) {
	tmp, err := template.ParseGlob(cfg.Template)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}
