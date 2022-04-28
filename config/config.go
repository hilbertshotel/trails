package config

type Mongo struct {
	Uri      string
	Database string
	Coll     string
}

type Config struct {
	HostAddr string
	Template string
	Mongo    Mongo
}

func Init() *Config {
	return &Config{
		HostAddr: "127.0.0.1:9990",
		Template: "templates/*",
		Mongo: Mongo{
			Uri:      "mongodb://localhost:27017",
			Database: "trails",
			Coll:     "workouts",
		},
	}
}
