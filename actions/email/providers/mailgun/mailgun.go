package mailgun

import (
	"bytes"
	"encoding/json"
	"example/vtr-mailer-service/structs"
	"example/vtr-mailer-service/tools"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

func Dispatch(mailgunConfig ServiceConfig, emailData structs.EmailData, htmlBody string) (MailgunResponse, error) {
	var data bytes.Buffer
	writer := multipart.NewWriter(&data)

	fields := map[string]string{
		"from":    emailData.From,
		"to":      strings.Join(emailData.To, ", "),
		"cc":      strings.Join(emailData.Cc, ", "),
		"bcc":     strings.Join(emailData.Bcc, ", "),
		"subject": emailData.Subject,
		"html":    htmlBody,
	}

	for field, value := range fields {
		if err := tools.AddFormField(writer, field, value); err != nil {
			return MailgunResponse{}, err
		}
	}

	if err := writer.Close(); err != nil {
		return MailgunResponse{}, err
	}

	req, err := createMailgunRequest(&data, writer, mailgunConfig)
	if err != nil {
		return MailgunResponse{}, err
	}

	return dispatchMailgunRequest(req)
}

func createMailgunRequest(data *bytes.Buffer, writer *multipart.Writer, mailgunConfig ServiceConfig) (*http.Request, error) {
	reqUrl := "https://api.mailgun.net/v3/" + mailgunConfig.Domain + "/messages"

	req, Error := http.NewRequest("POST", reqUrl, bytes.NewReader(data.Bytes()))
	if Error != nil {
		return nil, Error
	}
	req.SetBasicAuth(mailgunConfig.Username, mailgunConfig.Password)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	return req, nil
}

func dispatchMailgunRequest(req *http.Request) (DispatchMailgunResponse MailgunResponse, Error error) {
	client := &http.Client{
		Timeout: 10 * time.Second, // Exemplo de timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return MailgunResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return MailgunResponse{}, err
	}

	var mailResponse MailgunResponse
	if err := json.Unmarshal(body, &mailResponse); err != nil {
		return MailgunResponse{}, err
	}
	mailResponse.StatusCode = resp.StatusCode

	return mailResponse, nil
}
