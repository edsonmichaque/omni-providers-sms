package sms

import "io"

type Partner struct{}

type SendResponse struct{}

type SendRequest struct{}

type QueryStatusRequest struct{}

type QueryStatusResponse struct{}

type ReceiveResponse struct{}

type Parser interface {
	Parse(io.Reader) error
}

type Sender interface {
	Send(SendRequest) SendResponse
}

type Querier interface {
	QueryStatus(QueryStatusRequest) QueryStatusResponse
}

type Receiver interface {
	Receive() ReceiveResponse
}

type Provider interface {
	Parser
	Sender
	Querier
	Receiver
}

type Status struct {
	Code        StatusCode   `json:"code"`
	Reason      StatusReason `json:"reason"`
	Description string       `json:"description"`
}

type StatusCode string

type StatusReason string

const (
	StatusNotImplemented = StatusCode("NOT_IMPLEMENTED")
	StatusPending        = StatusCode("PENDING")
	StatusEnroute        = StatusCode("ENROUTE")
	StatusScheduled      = StatusCode("SCHEDULED")
	StatusDeleted        = StatusCode("DELETED")
	StatusAccepted       = StatusCode("ACCEPTED")
	StatusUnknown        = StatusCode("UNKNOWN")
	StatusSkipped        = StatusCode("SKIPPED")
	StatusUndeliverable  = StatusCode("UNDELIVERABLE")
	StatusDelivered      = StatusCode("DELIVERED")
	StatusExpired        = StatusCode("EXPIRED")
	StatusRejected       = StatusCode("REJECTED")
)

const (
	ReasonX = StatusReason("")
)

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
