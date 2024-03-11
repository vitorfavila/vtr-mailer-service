package config

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (c *Config) GetCors() gin.HandlerFunc {
	return cors.New(c.GetCorsConfig())
}

func (c *Config) GetCorsConfig() cors.Config {
	config := cors.DefaultConfig()

	if c.GetEnvironment() == "development" {
		fmt.Println("DEVELOPMENT CORS")
		config.AllowAllOrigins = true
		// config.AllowOrigins = []string{"*"}
	} else {
		config.AllowOrigins = c.GetAllowedOrigins()
	}

	return config
}

func (c *Config) GetAllowedOrigins() []string {
	return []string{
		"https://domain.com",
		"https://www.domain.com",
		"https://app.domain.com",
	}
}
