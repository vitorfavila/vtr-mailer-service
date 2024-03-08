package router

import (
	"example/vtr-mailer-service/application"
	"example/vtr-mailer-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupEmailTransactionGroup(app *application.Application, router *gin.RouterGroup) {
	router.POST("/create", func(c *gin.Context) {
		controllers.CreateEmailTransaction(app, c)
	})

	router.POST("/process/:transactionId", func(c *gin.Context) {
		controllers.ProcessEmailTransaction(app, c)
	})

	router.GET("/view/:transactionId", func(c *gin.Context) {
		controllers.ViewEmail(app, c)
	})

	router.GET("/get/:transactionId", func(c *gin.Context) {
		controllers.GetTransaction(app, c)
	})
}
