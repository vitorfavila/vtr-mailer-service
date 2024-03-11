package router

import (
	"example/vtr-mailer-service/application"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *application.Application) *gin.Engine {
	router := gin.New()
	if app.Cfg.GetEnvironment() == "development" {
		router.Use(gin.Logger())
	}
	router.Use(gin.Recovery())

	// CORS SETUP
	router.Use(app.Cfg.GetCors())

	emailTransactionGroup := router.Group("/email")
	{
		// emailTransactionGroup.Use(middleware.TokenAuthMiddleware())
		SetupEmailTransactionGroup(app, emailTransactionGroup)
	}

	router.GET("/health", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("OK"))
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return router
}
