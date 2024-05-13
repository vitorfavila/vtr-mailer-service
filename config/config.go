package config

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

type Config struct {
	// APP
	appEnvironment string
	SendMode       string

	// DB
	dbUser string
	dbPwd  string
	dbHost string
	dbPort string
	dbName string

	// Mailgun
	mailgunDomain string
	mailgunUser   string
	mailgunPwd    string
}

type DBConfig struct {
	User string
	Pwd  string
	Host string
	Port string
	Name string
}

func Get() *Config {
	conf := &Config{}
	// APP
	flag.StringVar(&conf.appEnvironment, "app_environment", os.Getenv("APP_ENVIRONMENT"), "Application environment")
	flag.StringVar(&conf.SendMode, "send_mode", os.Getenv("EMAIL_SEND_MODE"), "Test Mode for Email")

	// DB
	flag.StringVar(&conf.dbUser, "dbUser", os.Getenv("POSTGRES_USER"), "Database User")
	flag.StringVar(&conf.dbPwd, "dbPwd", os.Getenv("POSTGRES_PASSWORD"), "Database Password")
	flag.StringVar(&conf.dbHost, "dbHost", os.Getenv("POSTGRES_HOST"), "Database Host")
	flag.StringVar(&conf.dbPort, "dbPort", os.Getenv("POSTGRES_PORT"), "Database Port")
	flag.StringVar(&conf.dbName, "dbName", os.Getenv("POSTGRES_DB"), "Database Name")

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

func (c *Config) GetDBConnStr() string {
	user := url.PathEscape(c.dbUser)
	password := url.PathEscape(c.dbPwd)

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		c.dbHost,
		c.dbPort,
		c.dbName,
	)
}

func (c *Config) GetDBConfig() DBConfig {
	return DBConfig{
		User: c.dbUser,
		Pwd:  c.dbPwd,
		Host: c.dbHost,
		Port: c.dbPort,
		Name: c.dbName,
	}
}
