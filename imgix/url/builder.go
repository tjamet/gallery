package url

import (
	"net/url"
)

type formatter interface {
	URL(string, url.Values) string
}

type transformer interface {
	GetImgixURLValues() url.Values
}

type builder struct {
	values    url.Values
	formatter formatter

}

func With() *builder {
	return &builder{
		values: url.Values{},
	}
}

func (b *builder) Builder(formatter formatter) *builder {
	b.formatter = formatter
	return b
}

func (b *builder) Add(transformer transformer) *builder {
	for key, val := range(transformer.GetImgixURLValues()){
		b.values[key] = val
	}
	return b
}

func (b *builder) ForImage(path string) string {
	return b.formatter.URL(path, b.values)
}