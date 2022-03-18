package config

type Config struct {
	Port string `json:"port"`
}

func Get() *Config {
	return &Config{
		Port: "8888",
	}
}
