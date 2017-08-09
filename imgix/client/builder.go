package client

import (
	"github.com/tjamet/gallery/config"
	"github.com/parkr/imgix-go"
)

type builder struct {
	cfg config.Getter
}

type client struct {
	client imgix.Client
}

func With() *builder {
	return &builder{}
}

func (b *builder) Config(cfg config.Getter) *builder {
	b.cfg = cfg
	return b
}

func (b *builder) New() (*client) {
	domain, err := b.cfg.Get("domain")
	if err != nil {
		return nil
	}
	token, err := b.cfg.Get("signToken")
	if err != nil {
		return &client{imgix.NewClient(domain)}
	}
	return &client{imgix.NewClientWithToken(domain, token)}
}
