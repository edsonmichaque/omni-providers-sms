package smsprovider

import "io"

type ConfigReader interface {
	ReadConfig(io.Reader) error
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

type MessageProvider interface {
	ConfigReader
	MessageSender
	MessageQuerier
	MessageReceiver
	Scheduler
}
