package dep

// INITIALIZE CONFIG
func initConfig() *Config {
	return &Config{
		HostAddr: "0.0.0.0:9990",
		Template: "templates/*",
		ConnStr:  "user=postgres dbname=trails sslmode=disable host=/run/postgresql",
	}
}
