package smpp

import (
	"io"

	smsprovider "github.com/edsonmichaque/stencil-sms-providers"
)

type Provider struct{}

func (p Provider) Parse(r io.Reader) error {
	return nil
}

func (p Provider) Send(s smsprovider.SendRequest) smsprovider.SendResponse {
	return smsprovider.SendResponse{}
}

func (p Provider) QueryStatus(s smsprovider.QueryStatusRequest) smsprovider.QueryStatusResponse {
	return smsprovider.QueryStatusResponse{}
}
