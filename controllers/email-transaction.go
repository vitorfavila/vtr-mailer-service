package controllers

import (
	"example/vtr-mailer-service/actions/email"
	"example/vtr-mailer-service/actions/transaction"
	"example/vtr-mailer-service/application"
	"example/vtr-mailer-service/structs"
	"example/vtr-mailer-service/tools"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEmailTransaction(app *application.Application, c *gin.Context) {
	var newTransactionCreateRequest structs.TransactionCreateRequest

	if err := c.BindJSON(&newTransactionCreateRequest); err != nil {
		return
	}

	newTransaction := structs.Transaction{
		Id:                       tools.GenerateRandId(),
		Status:                   transaction.Created,
		TransactionCreateRequest: newTransactionCreateRequest,
	}

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
	tmpl, errTmpl := transaction.ParseGenerateTemplate(trans.TemplateID, tmplString, trans.Context)
	if errTmpl != nil {
		transaction.UpdateTransactionStatus(trans.Id, transaction.FailOnProcess)
		c.Status(http.StatusInternalServerError)
		return
	}

	trans = transaction.UpdateTransactionStatus(trans.Id, transaction.ProcessFinished)

	errSend := email.SendEmail(app, trans, tmpl)
	if errSend != nil {
		fmt.Println(errSend.Error())
		transaction.UpdateTransactionStatus(trans.Id, transaction.FailOnSend)
		c.Status(http.StatusInternalServerError)
		return
	}
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
	tmpl, errTmpl := transaction.ParseGenerateTemplate(trans.TemplateID, tmplString, trans.Context)
	if errTmpl != nil {
		transaction.UpdateTransactionStatus(trans.Id, transaction.FailOnProcess)
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
