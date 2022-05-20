package smpp

import (
	"io"

	smsprovider "github.com/edsonmichaque/stencil-providers-sms"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Provider struct {
	config Config
}

var _ smsprovider.MessageProvider = &Provider{}

func (p *Provider) ReadConfig(r io.Reader) error {
	return nil
}

func (p *Provider) SendMessages(s smsprovider.Request) smsprovider.Response {
	return smsprovider.Response{
		NotImplemented: smsprovider.Bool(false),
	}
}

func (p *Provider) QueryMessagesStatus(s smsprovider.Request) smsprovider.Response {
	return smsprovider.Response{
		NotImplemented: smsprovider.Bool(false),
	}
}

func (p *Provider) ScheduleMessages(string, smsprovider.Request) smsprovider.Response {
	return smsprovider.Response{
		NotImplemented: smsprovider.Bool(false),
	}
}

func (p *Provider) ReceiveMessages(s smsprovider.Request) smsprovider.Response {
	return smsprovider.Response{
		NotImplemented: smsprovider.Bool(false),
	}
}
