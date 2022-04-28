package config

type Config struct {
	HostAddr string
	ConnStr  string
	Template string
}

func Init() *Config {
	return &Config{
		HostAddr: "127.0.0.1:9990",
		ConnStr:  "user=postgres dbname=trails sslmode=disable host=/run/postgresql",
		Template: "templates/*",
	}
}
