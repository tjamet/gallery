package client

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/url"
	"fmt"
)

type testConfig struct {
	cfgs map[string]string
}

func (t *testConfig) Get(name string) (string, error) {
	v, ok := t.cfgs[name]
	if ok {
		return v, nil
	}
	return "", fmt.Errorf("Key does not exist")
}

func TestBuilder_New(t *testing.T) {
	var b *builder
	b = With().Config(&testConfig{
		cfgs: map[string]string {"domain":"test.imgix.net"},
	})
	cfg := b.New()
	assert.NotNil(t, cfg)
	assert.Equal(t, "test.imgix.net", cfg.client.Hosts(0))
	assert.True(t, cfg.client.Secure())
	assert.Equal(t, "https://test.imgix.net/test.jpg", cfg.URL("test.jpg", nil))
	values := url.Values{"w": []string{"10"}}
	assert.Equal(t, "https://test.imgix.net/test.jpg?w=10", cfg.URL("test.jpg", values))

	b = With().Config(&testConfig{
		cfgs: map[string]string {"domain":"test.imgix.net", "signToken": "testToken"},
	})
	cfg = b.New()
	assert.NotNil(t, cfg)
	assert.Equal(t, "test.imgix.net", cfg.client.Hosts(0))
	assert.Equal(t, "https://test.imgix.net/test.jpg?s=91eef7ef61c26477e3ef09459ea53cb2", cfg.URL("test.jpg", nil))
	assert.True(t, cfg.client.Secure())
}
