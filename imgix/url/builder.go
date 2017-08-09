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

func merge(ins ...url.Values) url.Values {
	r := url.Values{}

	for _, in := range ins {
		for key, val := range in {
			r[key] = val
		}
	}
	return r
}

func (b *builder) Clone() *builder {
	return &builder{
		formatter: b.formatter,
		values:    merge(b.values),
	}
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

func (b *builder) Update(values ...url.Values) *builder {
	values = append([]url.Values{b.values}, values...)
	b.values = merge(values...)
	return b
}

func (b *builder) Add(transformer transformer) *builder {
	return b.Update(transformer.GetImgixURLValues())
}

func (b *builder) ForImage(path string) string {
	return b.formatter.URL(path, b.values)
}
