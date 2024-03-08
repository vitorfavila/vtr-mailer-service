package structs

type TransactionStatus int

type Transaction struct {
	Id     int64             `json:"id,omitempty" db:"id"`
	Status TransactionStatus `json:"status" db:"status"`
	TransactionCreateRequest
}

type EmailData struct {
	To       []string
	Cc       []string
	Bcc      []string
	From     string
	Subject  string
	HtmlBody string
	SendMode string
}
