package smsprovider

type Partner struct{}

type SendMessageResponse struct{}

type SendMessageRequest struct{}

type QueryMessageStatusRequest struct{}

type QueryMessageStatusResponse struct{}

type ReceiveMessageResponse struct{}

type Status struct {
	Code        StatusCode   `json:"code"`
	Reason      StatusReason `json:"reason"`
	Description string       `json:"description"`
}

type Destination struct {
	To string `json:"to"`
}

type MessageStats struct {
	Count string `json:"count"`
}

type Quotation struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Message struct {
	ID           string        `json:"id,omitempty"`
	Text         string        `json:"text,omitempty"`
	From         string        `json:"from,omitempty"`
	SubmittedAt  string        `json:"submitted_at,omitempty"`
	DoneAt       string        `json:"done_at,omitempty"`
	Destinations []Destination `json:"destinations,omitempty"`
	Status       Status        `json:"status,omitempty"`
	Stats        MessageStats  `json:"stats,omitempty"`
	Quotation    Quotation     `json:"quotation,omitempty"`
}
