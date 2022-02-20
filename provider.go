package example

import (
	"gitlab.com/edsonmichaque/libpx"
	"gitlab.com/edsonmichaque/libpx/card"
	"gitlab.com/edsonmichaque/libpx/currency"
	"gitlab.com/edsonmichaque/libpx/schema"
)

func New() Provider {
	return Provider{}
}

type Provider struct {
	Gateway
}

func (p Provider) Schema() map[string]schema.Schema {
	return map[string]schema.Schema{
		"api_key": {
			Type:     schema.String,
			Required: true,
		},
	}
}

func (p Provider) Currencies() []currency.Currency {
	return []currency.Currency{
		currency.USD,
		currency.MZN,
	}
}

func (p Provider) SupportedSources() []string {
	return []string{
		"bank",
	}
}

func (p Provider) Configure(args ...libpx.ConfigOption) error {

	return nil
}

type Gateway struct{}

func (g Gateway) Authorize(card card.Card, amount libpx.Amount, opts ...libpx.Option) (*libpx.Authorization, error) {
	options := libpx.Options{}

	for _, opt := range opts {
		opt(&options)
	}

	return nil, libpx.NotSupportedError{}
}

func (g Gateway) Capture(auth libpx.Authorization, amount libpx.Amount, opts ...libpx.Option) (*libpx.Capture, error) {
	options := libpx.Options{}

	for _, opt := range opts {
		opt(&options)
	}

	return nil, libpx.NotSupportedError{}
}

func (g Gateway) Purchase(auth libpx.Authorization, amount libpx.Amount, opts ...libpx.Option) (*libpx.Purchase, error) {
	options := libpx.Options{}

	for _, opt := range opts {
		opt(&options)
	}

	return nil, libpx.NotSupportedError{}
}

func (g Gateway) Refund(auth libpx.Authorization, amount libpx.Amount) (*libpx.Refund, error) {
	return nil, libpx.NotSupportedError{}
}

func (g Gateway) Void(auth libpx.Authorization) (*libpx.Void, error) {
	return nil, libpx.NotSupportedError{}
}

func (g Gateway) Verify(card card.Card, opts ...libpx.Option) (*libpx.Verification, error) {
	options := libpx.Options{}

	for _, opt := range opts {
		opt(&options)
	}

	return nil, libpx.NotSupportedError{}
}
