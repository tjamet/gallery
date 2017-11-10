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
	if kwdsObj, ok := m.Metadata["Keywords"]; ok {
		newKwds := []string{}
		if kwds, ok := kwdsObj.([]interface{}); ok {
			for _, key := range kwds {
				if keep, ok := f.Keywords[strings.ToLower(key.(string))]; ok {
					if keep {
						newKwds = append(newKwds, key.(string))
					}
				} else if f.Default {
					newKwds = append(newKwds, key.(string))
				}
			}
		} else if v, ok := kwdsObj.(string); ok {
			if keep, ok := f.Keywords[strings.ToLower(v)]; ok {
				if keep {
					newKwds = append(newKwds, v)
				}
			} else if f.Default {
				newKwds = append(newKwds, v)
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
