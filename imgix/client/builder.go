package client

import (
	"github.com/tjamet/gallery/config"
	"github.com/parkr/imgix-go"
)

type builder struct {
	cfg config.Getter
}

func With() *builder {
	return &builder{}
}

func (b *builder) Config(cfg config.Getter) *builder {
	b.cfg = cfg
	return b
}

func (b *builder) New() (*imgix.Client) {
	domain, err := b.cfg.Get("domain")
	if err != nil {
		return nil
	}
	token, err := b.cfg.Get("signToken")
	var client imgix.Client
	if err != nil {
		client = imgix.NewClient(domain)
	} else {
		client = imgix.NewClientWithToken(domain, token)
	}
	return &client
}