package mailgun

type MailgunResponse struct {
	ID         string `json:"id"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type ServiceConfig struct {
	Domain   string
	Username string
	Password string
}
