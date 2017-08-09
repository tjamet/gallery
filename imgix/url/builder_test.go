package url

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testTransformer struct{}

func (t testTransformer) GetImgixURLValues() url.Values {
	return url.Values{
		"w": {"100"},
		"h": {"100"},
	}
}

type testTransformer2 struct{}

func (t testTransformer2) GetImgixURLValues() url.Values {
	return url.Values{
		"w": {"140"},
		"g": {"1"},
	}
}

type testFormatter struct{}

func (t testFormatter) URL(path string, values url.Values) string {
	return path + "?" + values.Encode()
}

func TestBuilder_Add(t *testing.T) {
	b := With().Add(testTransformer{})
	assert.Equal(t, url.Values{
		"w": {"100"},
		"h": {"100"},
	}, b.values)
	b.Add(testTransformer2{})
	assert.Equal(t, url.Values{
		"w": {"140"},
		"h": {"100"},
		"g": {"1"},
	}, b.values)
}
func TestBuilder_Update(t *testing.T) {
	b := With().Update(
		url.Values{
			"w": {"100"},
			"h": {"100"},
		},
		url.Values{
			"w": {"140"},
			"g": {"1"},
		},
	)
	assert.Equal(t, url.Values{
		"w": {"140"},
		"h": {"100"},
		"g": {"1"},
	}, b.values)
}

func TestBuilder_ForImage(t *testing.T) {
	b := With().Builder(testFormatter{})
	b.values.Add("w", "12")
	assert.Equal(t, "test.jpg?w=12", b.ForImage("test.jpg"))
}
