package smpp

import (
	"io"

	"github.com/edsonmichaque/stencil/sms"
)

type Provider struct{}

func (p Provider) Parse(r io.Reader) error {
	return nil
}

func (p Provider) Send(s sms.SendRequest) sms.SendResponse {
	return sms.SendResponse{}
}

func (p Provider) QueryStatus(s sms.QueryStatusRequest) sms.QueryStatusResponse {
	return sms.QueryStatusResponse{}
}
