package smsprovider

import "io"

type Proxy struct {
	Provider Provider
}

func (p Proxy) NewProvider(build ProviderBuilder, r io.Reader) (Provider, error) {
	newProvider := build()

	if err := newProvider.ParseConfig(r); err != nil {
		return nil, err
	}

	return newProvider, nil
}

func (p Proxy) ScheduleMessages(provider Provider, date string, req Request) Response {
	if err := p.validateRequest(req); err != nil {
		return Response{}
	}

	resp := provider.ScheduleMessages(date, req)
	if err := p.validateResponse(resp); err != nil {
		return Response{}
	}

	return resp
}

func (p Proxy) SendMessages(provider Provider, req Request) Response {
	if err := p.validateRequest(req); err != nil {
		return Response{}
	}

	resp := provider.SendMessages(req)
	if err := p.validateResponse(resp); err != nil {
		return Response{}
	}

	return resp
}

func (p Proxy) QueryMessagesStatus(provider Provider, req Request) Response {
	if err := p.validateRequest(req); err != nil {
		return Response{}
	}

	resp := provider.QueryMessagesStatus(req)
	if err := p.validateResponse(resp); err != nil {
		return Response{}
	}

	return resp
}

func (p Proxy) ReceiveMessages(provider Provider, req Request) Response {
	if err := p.validateRequest(req); err != nil {
		return Response{}
	}

	resp := provider.ReceiveMessages(req)
	if err := p.validateResponse(resp); err != nil {
		return Response{}
	}

	return resp
}

func (p Proxy) validateRequest(r Request) error {
	return nil
}

func (p Proxy) validateResponse(r Response) error {
	return nil
}
