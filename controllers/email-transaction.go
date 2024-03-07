package controllers

import (
	"bytes"
	"example/vtr-mailer-service/actions/transaction"
	"example/vtr-mailer-service/application"
	"example/vtr-mailer-service/structs"
	"example/vtr-mailer-service/tools"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

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

	tmplString := app.Template.Templates.Get(trans.TemplateID)
	tmpl, err := template.New(trans.TemplateID).Parse(tmplString)
	if err != nil {
		fmt.Println("1 " + err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	var htmlBuffer bytes.Buffer
	err = tmpl.Execute(&htmlBuffer, trans.Context)
	if err != nil {
		fmt.Println("2 " + err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	fmt.Println(htmlBuffer.String())

	trans.Status = transaction.ProcessPending
	time.Sleep(time.Second * 2)
	trans.Status = transaction.ProcessFinished
	transaction.UpdateTransaction(trans)

	c.IndentedJSON(http.StatusOK, trans)
}
