package controllers

import (
	"example/vtr-mailer-service/actions/email"
	"example/vtr-mailer-service/actions/transaction"
	"example/vtr-mailer-service/application"
	"example/vtr-mailer-service/structs"
	"example/vtr-mailer-service/tools"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEmailTransaction(app *application.Application, c *gin.Context) {
	var newTransactionCreateRequest structs.TransactionCreateRequest

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newTransactionCreateRequest); err != nil {
		return
	}

	newTransaction := structs.Transaction{
		Id:                       tools.GenerateRandId(),
		Status:                   transaction.Created,
		TransactionCreateRequest: newTransactionCreateRequest,
	}

	// Add the new album to the slice.
	transaction.AddTransaction(newTransaction)
	c.IndentedJSON(http.StatusCreated, newTransaction)
}

func ProcessEmailTransaction(app *application.Application, c *gin.Context) {
	transactionId, _ := strconv.ParseInt(c.Param("transactionId"), 10, 64)
	trans, err := transaction.GetTransaction(transactionId)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	trans = transaction.UpdateTransactionStatus(trans.Id, transaction.ProcessPending)

	tmplString := app.Template.Templates.Get(trans.TemplateID)
	tmplParsed, errParsing := transaction.ParseTemplate(trans.TemplateID, tmplString)
	if errParsing != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	tmpl, errTmpl := transaction.GenerateTemplate(*tmplParsed, trans.Context)
	if errTmpl != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	trans = transaction.UpdateTransactionStatus(trans.Id, transaction.ProcessFinished)

	email.SendEmail(app, trans, tmpl)
	trans = transaction.UpdateTransactionStatus(trans.Id, transaction.Sent)

	c.IndentedJSON(http.StatusOK, trans)
}

func ViewEmail(app *application.Application, c *gin.Context) {
	transactionId, _ := strconv.ParseInt(c.Param("transactionId"), 10, 64)

	trans, err := transaction.GetTransaction(transactionId)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	tmplString := app.Template.Templates.Get(trans.TemplateID)
	tmplParsed, errParsing := transaction.ParseTemplate(trans.TemplateID, tmplString)
	if errParsing != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	tmpl, errTmpl := transaction.GenerateTemplate(*tmplParsed, trans.Context)
	if errTmpl != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Data(200, "text/html; charset=utf-8", []byte(tmpl))
}

func GetTransaction(app *application.Application, c *gin.Context) {
	transactionId, _ := strconv.ParseInt(c.Param("transactionId"), 10, 64)

	trans, err := transaction.GetTransaction(transactionId)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, trans)
}
