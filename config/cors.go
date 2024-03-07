package config

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (c *Config) GetCors() gin.HandlerFunc {
	return cors.New(c.GetCorsOpenConfig())
}

func (c *Config) GetCorsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AddAllowHeaders("ELQ-TKA")
	config.AllowOrigins = c.GetAllowedOrigins()

	if c.GetEnvironment() == "development" {
		fmt.Println("DEVELOPMENT CORS")
		config.AllowOrigins = []string{"*"}
	}

	return config
}

func (c *Config) GetCorsOpenConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AddAllowHeaders("ELQ-TKA", "Authorization")
	config.AllowAllOrigins = true

	return config
}

func (c *Config) GetAllowedOrigins() []string {
	return []string{
		"https://elquarto.com",
		"https://www.elquarto.com",
		"https://elquarto.com/blog",
		"https://www.elquarto.com/blog",
	}
}
