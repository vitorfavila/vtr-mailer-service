package config

import (
	"flag"
	"os"
)

type Config struct {
	// APP
	appEnvironment string

	// DB
	// dbUser string
	// dbPwd  string
	// dbHost string
	// dbPort string
	// dbName string
}

func Get() *Config {
	conf := &Config{}
	// APP
	flag.StringVar(&conf.appEnvironment, "app_environment", os.Getenv("APP_ENVIRONMENT"), "Application environment")

	flag.Parse()

	return conf
}

func (c *Config) GetEnvironment() string {
	return c.appEnvironment
}
