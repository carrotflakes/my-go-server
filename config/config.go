package config

import "os"

type Config struct {
	Env  string `env:"env"`
	Port string `json:"port"`
}

func Get() *Config {
	return &Config{
		Env:  os.Getenv("GO_ENV"),
		Port: "8888",
	}
}
