package mailgun

import (
	"bytes"
	"example/vtr-mailer-service/structs"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func SendEmail(domain string, username string, password string, emailData structs.EmailData, htmlBody string) (MailgunResponse string, Error error) {
	var err error

	domainName := domain
	reqUrl := "https://api.mailgun.net/v3/" + domainName + "/messages"

	data := &bytes.Buffer{}
	writer := multipart.NewWriter(data)
	fromFw, _ := writer.CreateFormField("from")
	_, err = io.Copy(fromFw, strings.NewReader(emailData.From))
	if err != nil {
		return "", err
	}
	toFw, _ := writer.CreateFormField("to")
	_, err = io.Copy(toFw, strings.NewReader(parseMultipleEmailsToString(emailData.To)))
	if err != nil {
		return "", err
	}
	subjectFw, _ := writer.CreateFormField("subject")
	_, err = io.Copy(subjectFw, strings.NewReader(emailData.Subject))
	if err != nil {
		return "", err
	}
	htmlFw, _ := writer.CreateFormField("html")
	_, err = io.Copy(htmlFw, strings.NewReader(htmlBody))
	if err != nil {
		return "", err
	}
	writer.Close()

	payload := bytes.NewReader(data.Bytes())
	req, err := http.NewRequest("POST", reqUrl, payload)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(username, password)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// DEBUG
	fmt.Println(res.StatusCode)
	// fmt.Println(string(body))

	return string(body), nil
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
