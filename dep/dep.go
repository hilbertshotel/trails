package dep

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"runtime"

	_ "github.com/lib/pq"
)

// LOAD DEPENDENCIES
func Load(path string) (*Dependencies, error) {
	log := initLogger()
	cfg, err := loadConfig(path)
	if err != nil {
		return nil, err
	}
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

// LOAD CONFIG
func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
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
	tmp, err := template.ParseGlob(cfg.TmpDir)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}
