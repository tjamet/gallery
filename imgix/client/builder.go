package client

import (
	"github.com/parkr/imgix-go"
	"github.com/tjamet/gallery/config"
)

type builder struct {
	cfg config.Getter
}

type client struct {
	client imgix.Client
}

func With() *builder {
	return &builder{
		cfg: config.With().
			Path("~/.imgix/credentials").
			Section("default").
			Build(),
	}
}

func (b *builder) Config(cfg config.Getter) *builder {
	b.cfg = cfg
	return b
}

func (b *builder) New() *client {
	domain, err := b.cfg.Get("domain")
	if err != nil {
		return nil
	}
	token, err := b.cfg.Get("signToken")
	if err != nil {
		return &client{imgix.NewClient(domain + ".imgix.net")}
	}
	return &client{imgix.NewClientWithToken(domain+".imgix.net", token)}
}
