package filters

import (
	"github.com/tjamet/gallery/metadata"
	"strings"
)

type KeyWordFilter struct {
	Keywords map[string]bool
	Default  bool
}

func (f *KeyWordFilter) Filter(m *metadata.Metadata) *metadata.Metadata {
	var output metadata.Metadata
	output = *m
	if kwds, ok := m.Metadata["Keywords"].([]interface{}); ok {
		newKwds := []string{}
		for _, key := range kwds {
			if keep, ok := f.Keywords[strings.ToLower(key.(string))]; ok {
				if keep {
					newKwds = append(newKwds, key.(string))
				}
			} else if f.Default {
				newKwds = append(newKwds, key.(string))
			}
		}
		output.Metadata["Keywords"] = newKwds
	}
	return &output
}

func MetadataFilterFromWhitelist(list []string) *KeyWordFilter {
	m := make(map[string]bool)
	for _, key := range (list) {
		strings.ToLower(key)
		m[strings.ToLower(key)] = true
	}
	return &KeyWordFilter{
		Keywords: m,
		Default:  false,
	}
}
