package metadata

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMetadata_GetMetadata(t *testing.T) {
	m := FromFile("../resources/images/test.jpg")
	expected := map[string]interface{}{
		"Artist":    "Thibault JAMET",
		"Copyright": "Photo by Thibault JAMET - www.thibaultjamet.fr",
		"Date":      "2014:04:27 02:15:54",
		"ID":        "B359240552455B49FFB7D34475CD0C81",
		"Title":     "lights",
		"Creator":   "Thibault JAMET",
	}
	metadata, err := m.GetMetadata()
	assert.NoError(t, err)

	keywords, ok := metadata["Keywords"]
	assert.True(t, ok, "metadata should have keywords")
	delete(metadata, "Keywords")
	assert.Contains(t, keywords, "Jay Style")
	assert.Contains(t, keywords, "Le Loft")
	assert.Contains(t, keywords, "DJ")
	assert.Equal(t, expected, metadata)
}
