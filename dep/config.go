package dep

// INITIALIZE CONFIG
func initConfig() *Config {
	return &Config{
		HostAddr: "127.0.0.1:9990",
		Template: "templates/*",
		Mongo: Mongo{
			Uri:        "mongodb://localhost:27017",
			Database:   "trails",
			Collection: "workouts",
		},
	}
}
