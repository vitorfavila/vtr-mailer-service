package tools

import "mime/multipart"

// addFormField adiciona um campo ao writer de um formulário multipart.
func AddFormField(writer *multipart.Writer, fieldName, value string) error {
	part, err := writer.CreateFormField(fieldName)
	if err != nil {
		return err
	}
	_, err = part.Write([]byte(value))
	if err != nil {
		return err
	}
	return nil
}
