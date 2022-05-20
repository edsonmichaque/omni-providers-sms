package smpp

import (
	"io"

	smsprovider "github.com/edsonmichaque/stencil-providers-sms"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ProviderOption func(*Provider)

type Provider struct {
	config Config
}

func New() smsprovider.Provider {
	return &Provider{}
}

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

func (p *Provider) SetLogger(l smsprovider.Logger) {

}
