package transaction

import (
	"bytes"
	"example/vtr-mailer-service/structs"
	"text/template"
)

func ParseTemplate(templateId string, htmlTemplate string) (*template.Template, error) {
	tmpl, err := template.New(templateId).Parse(htmlTemplate)
	if err != nil {
		return tmpl, err
	}

	return tmpl, nil
}

func GenerateTemplate(tmpl template.Template, context structs.Context) (string, error) {
	var htmlBuffer bytes.Buffer
	err := tmpl.Execute(&htmlBuffer, context)
	if err != nil {
		return "", err
	}

	return htmlBuffer.String(), nil
}

func ParseGenerateTemplate(templateId string, htmlTemplate string, context structs.Context) (string, error) {
	var err error
	tmplParsed, err := ParseTemplate(templateId, htmlTemplate)
	if err != nil {
		return "", err
	}

	tmpl, err := GenerateTemplate(*tmplParsed, context)
	if err != nil {
		return "", err
	}

	return tmpl, nil
}
