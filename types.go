package smsprovider

type Partner struct{}

func Bool(b bool) *bool {
	return &b
}

type Response struct {
	Error    *Error          `json:"error"`
	Messages []MessageStatus `json:"messages,omitempty"`
}

type Error struct {
	Code ErrorCode `json:"code,omitempty"`
}

type Request struct {
	Messages []Message `json:"messages,omitempty"`
	Schedule *string   `json:"schedule"`
}

type Status struct {
	Code        StatusCode `json:"code"`
	Reason      ReasonCode `json:"reason"`
	Description string     `json:"description"`
}

type Destination struct {
	To string `json:"to"`
}

type Quotation struct {
	Amount   string      `json:"amount"`
	Currency string      `json:"currency"`
	Count    int         `json:"count"`
	Size     MessageSize `json:"size"`
}

type MessageSize struct {
	Characters int `json:"characters"`
	Bytes      int `json:"bytes"`
}

type MessageStatus struct {
	MessageID         string `json:"message_id,omitempty"`
	ProviderMessageID string `json:"provider_message_id,omitempty"`
	SubmittedAt       string `json:"submitted_at,omitempty"`
	DoneAt            string `json:"done_at,omitempty"`
	Status            Status `json:"status,omitempty"`
}

type Message struct {
	ID           string        `json:"id,omitempty"`
	ProviderID   string        `json:"provider_id,omitempty"`
	Text         string        `json:"text,omitempty"`
	From         string        `json:"from,omitempty"`
	SubmittedAt  string        `json:"submitted_at,omitempty"`
	DoneAt       string        `json:"done_at,omitempty"`
	Destinations []Destination `json:"destinations,omitempty"`
	Status       Status        `json:"status,omitempty"`
	Quotation    Quotation     `json:"quotation,omitempty"`
}
