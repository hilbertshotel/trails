package dep

// INITIALIZE CONFIG
func initConfig() *Config {
	return &Config{
		HostAddr: "127.0.0.1:9990",
		Template: "templates/*",
		ConnStr:  "user=postgres dbname=uncle password=postgres sslmode=disable",
	}
}
