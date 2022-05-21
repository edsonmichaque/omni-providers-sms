package smsprovider

import "io"

type ProviderBuilder func() Provider

type ConfigReader interface {
	ParseConfig(io.Reader) error
}

type Scheduler interface {
	ScheduleMessages(string, Request) Response
}

type MessageSender interface {
	SendMessages(Request) Response
}

type MessageQuerier interface {
	QueryMessagesStatus(Request) Response
}

type MessageReceiver interface {
	ReceiveMessages(Request) Response
}

type Capabilities interface {
	Capabilities() []Capability
}

type Info interface {
	ID() string
	Name() string
	Version() string
}

type Capability string

const (
	CapabilitySendMessage         = Capability("SEND_MESSAGE")
	CapabilityScheduleMessages    = Capability("SCHEDULE_MESSAGES")
	CapabilityQueryMessagesStatus = Capability("QUERY_MESSAGES")
	CapabilityReceiveMessages     = Capability("RECEIVE_MESSAGES")
)

type Provider interface {
	ConfigReader
	MessageSender
	MessageQuerier
	MessageReceiver
	Scheduler
	Capabilities
}
