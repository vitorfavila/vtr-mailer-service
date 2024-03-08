package email

import (
	"bytes"
	"example/vtr-mailer-service/application"
	"example/vtr-mailer-service/structs"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func SendEmail(app *application.Application, transaction structs.Transaction, htmlBody string) {
	var err error
	domain, username, password := app.Cfg.GetMailgunConfig()

	domainName := domain
	reqUrl := "https://api.mailgun.net/v3/" + domainName + "/messages"

	data := &bytes.Buffer{}
	writer := multipart.NewWriter(data)
	fromFw, _ := writer.CreateFormField("from")
	_, err = io.Copy(fromFw, strings.NewReader(transaction.From))
	if err != nil {
		panic(err)
	}
	toFw, _ := writer.CreateFormField("to")
	_, err = io.Copy(toFw, strings.NewReader(parseMultipleEmailsToString(transaction.To)))
	if err != nil {
		panic(err)
	}
	subjectFw, _ := writer.CreateFormField("subject")
	_, err = io.Copy(subjectFw, strings.NewReader(transaction.Subject))
	if err != nil {
		panic(err)
	}
	htmlFw, _ := writer.CreateFormField("html")
	_, err = io.Copy(htmlFw, strings.NewReader(htmlBody))
	if err != nil {
		panic(err)
	}
	writer.Close()

	payload := bytes.NewReader(data.Bytes())
	req, err := http.NewRequest("POST", reqUrl, payload)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(username, password)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.StatusCode)
	fmt.Println(string(body))
}

func parseMultipleEmailsToString(emails []string) string {
	emailString := ""

	for _, email := range emails {
		if len(emailString) == 0 {
			emailString = email
			continue
		}

		emailString = emailString + ", " + email
	}

	return emailString
}
