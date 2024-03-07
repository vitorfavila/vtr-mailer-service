package structs

import (
	"time"
)

type TransactionCreateRequest struct {
	TemplateID   string    `json:"template_id" db:"template_id"`
	DeliveryTime time.Time `json:"delivery_time" db:"delivery_time"`
	Subject      string    `json:"subject" db:"subject"`
	From         string    `json:"from" db:"from"`
	To           []string  `json:"to" db:"to"`
	Cc           []string  `json:"cc" db:"cc"`
	Bcc          []string  `json:"bcc" db:"bcc"`
	Context      Context   `json:"context" db:"context"`
	Testmode     bool      `json:"testmode" db:"testmode"`
	Dryrun       bool      `json:"dryrun" db:"dryrun"`
	Tags         []string  `json:"tags" db:"tags"`
}

type Context struct {
	Locator    string    `json:"locator" db:"locator"`
	Etkt       string    `json:"etkt" db:"etkt"`
	Passengers []string  `json:"passengers" db:"passengers"`
	Journey    []Journey `json:"journey" db:"journey"`
	Notes      []Note    `json:"notes" db:"notes"`
}

type Journey struct {
	Iata    string `json:"iata" db:"iata"`
	Airport string `json:"airport" db:"airport"`
	City    string `json:"city" db:"city"`
	Country string `json:"country" db:"country"`
}

type Note struct {
	Title string `json:"title" db:"title"`
	Text  string `json:"text" db:"text"`
}
