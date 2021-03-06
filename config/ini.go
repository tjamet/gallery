package config

import (
	"log"
	user "os/user"

	"github.com/go-ini/ini"
)

type INIConfig struct {
	path    string
	section string
	cfg     *ini.File
}

func With() *INIConfig {
	c := &INIConfig{}
	return c.Section("default")
}

func (c *INIConfig) Section(name string) *INIConfig {
	c.section = name
	return c
}

func (c *INIConfig) Path(path string) *INIConfig {
	if path[:2] == "~/" {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		path = usr.HomeDir + "/" + path[2:]
	}
	c.path = path
	return c
}

func (c *INIConfig) Build() *INIConfig {
	cfg, err := ini.Load(c.path)
	if err != nil {
		log.Fatal(err)
	}
	c.cfg = cfg
	return c
}

func (c *INIConfig) Get(name string) (string, error) {
	if c.cfg == nil {
		c.Build()
	}
	s, err := c.cfg.GetSection(c.section)
	if err != nil {
		return "", err
	}
	k, err := s.GetKey(name)
	if err != nil {
		return "", err
	}
	return k.Value(), nil
}
