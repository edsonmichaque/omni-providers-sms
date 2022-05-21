package smpp50

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

func New() *Provider {
	return &Provider{}
}

func (p *Provider) ParseConfig(r io.Reader) error {
	return nil
}

func (p *Provider) SendMessages(s smsprovider.Request) smsprovider.Response {
	return smsprovider.ResponseNotImplemented
}

func (p *Provider) QueryMessagesStatus(s smsprovider.Request) smsprovider.Response {
	return smsprovider.ResponseNotImplemented
}

func (p *Provider) ScheduleMessages(string, smsprovider.Request) smsprovider.Response {
	return smsprovider.ResponseNotImplemented
}

func (p *Provider) ReceiveMessages(s smsprovider.Request) smsprovider.Response {
	return smsprovider.ResponseNotImplemented
}

func (p *Provider) Capabilities() []smsprovider.Capability {
	return []smsprovider.Capability{
		smsprovider.CapabilityQueryMessagesStatus,
	}
}
