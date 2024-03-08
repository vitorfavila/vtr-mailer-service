package config

import (
	"flag"
	"os"
)

type Config struct {
	// APP
	appEnvironment string
	SendMode       string

	// DB
	// dbUser string
	// dbPwd  string
	// dbHost string
	// dbPort string
	// dbName string

	// Mailgun
	mailgunDomain string
	mailgunUser   string
	mailgunPwd    string
}

func Get() *Config {
	conf := &Config{}
	// APP
	flag.StringVar(&conf.appEnvironment, "app_environment", os.Getenv("APP_ENVIRONMENT"), "Application environment")
	flag.StringVar(&conf.SendMode, "send_mode", os.Getenv("EMAIL_SEND_MODE"), "Test Mode for Email")

	// Mailgun
	flag.StringVar(&conf.mailgunDomain, "mailgunDomain", os.Getenv("MAILGUN_DOMAIN_NAME"), "Mailgun Domain Name")
	flag.StringVar(&conf.mailgunUser, "mailgunUser", os.Getenv("MAILGUN_AUTH_LOGIN"), "Mailgun Auth Login")
	flag.StringVar(&conf.mailgunPwd, "mailgunPwd", os.Getenv("MAILGUN_AUTH_PWD"), "Mailgun Auth Password")

	flag.Parse()

	return conf
}

func (c *Config) GetEnvironment() string {
	return c.appEnvironment
}

func (c *Config) GetMailgunConfig() (domain, username, password string) {
	return c.mailgunDomain, c.mailgunUser, c.mailgunPwd
}
