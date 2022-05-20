package smpp

import (
	"encoding/json"
	"io"

	smsprovider "github.com/edsonmichaque/stencil-sms-providers"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Provider struct {
	config Config
}

func (p *Provider) Parse(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &p.config); err != nil {
		return err
	}

	return nil
}

func (p *Provider) Send(s smsprovider.SendMessageRequest) smsprovider.SendMessageResponse {
	return smsprovider.SendMessageResponse{}
}

func (p *Provider) QueryStatus(s smsprovider.QueryMessageStatusRequest) smsprovider.QueryMessageStatusResponse {
	return smsprovider.QueryMessageStatusResponse{}
}
