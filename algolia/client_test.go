package algolia

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"time"
)

func TestClient(t *testing.T) {
	indexName := "test-" + uuid.New().String()
	builder := With().Index(indexName)
	client, err := builder.newAdminClient()
	assert.NoError(t, err)
	indexes, err := client.ListIndexes()
	assert.NoError(t, err)
	originalLength := len(indexes)

	object := algoliasearch.Object{
		"testKey" : "value",
		"objectID": uuid.New().String(),
	}
	err = builder.UpdateObjects(object)
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)
	client, err = builder.newAdminClient()
	assert.NoError(t, err)
	indexes, err = client.ListIndexes()
	assert.NoError(t, err)
	newLength := len(indexes)
	assert.Equal(t, originalLength+1, newLength)

	index := client.InitIndex(indexName)
	res, err := index.Search("value", nil)
	assert.NoError(t, err)
	assert.Equal(t,1, res.NbHits)

	client.DeleteIndex(indexName)
}

