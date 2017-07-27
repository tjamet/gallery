package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os/user"
)

func TestHomeDir(t *testing.T) {
	ini := With().Path("~/some.path")
	usr, err := user.Current()
	assert.NoError(t, err)
	assert.Equal(t, usr.HomeDir + "/some.path", ini.path)
}

func TestGet(t *testing.T){
	value, err := With().
		Path("resources/testConfig").
		Build().
		Get("myKey")
	assert.NoError(t, err)
	assert.Equal(t, "my value", value)
}

func TestNotBuilt(t *testing.T){
	value, err := With().
		Path("resources/testConfig").
		Get("myKey")
	assert.NoError(t, err)
	assert.NotNil(t, value)
}