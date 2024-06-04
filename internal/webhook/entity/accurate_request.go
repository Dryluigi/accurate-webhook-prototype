package entity

type AccuratePayload struct {
	DatabaseID int64     `json:"databaseId"`
	Type       string    `json:"type"`
	Timestamp  string    `json:"timestamp"`
	UUID       string    `json:"uuid"`
	Data       []DataObj `json:"data"`
}

type DataObj struct {
	SalesInvoiceID          *int    `json:"salesInvoiceId"`
	SalesInvoiceNo          *string `json:"salesInvoiceNo"`
	SalesInvoiceTotalAmount *string `json:"salesInvoiceTotalAmount"`
	Action                  string  `json:"action"`
}
