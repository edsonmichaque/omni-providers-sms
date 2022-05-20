package smsprovider

import "io"

type ConfigParser interface {
	ParseConfig(io.Reader) error
}

type MessageSender interface {
	SendMessage(SendMessageRequest) SendMessageResponse
}

type MessageQuerier interface {
	QueryMessageStatus(QueryMessageStatusRequest) QueryMessageStatusResponse
}

type MessageReceiver interface {
	ReceiveMessages() ReceiveMessageResponse
}

type MessageProvider interface {
	ConfigParser
	MessageSender
	MessageQuerier
	MessageReceiver
}
