package algolia

import (
	"os"
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	if os.Getenv("GALLERY_INTEGRATION") == "" {
		t.Skip("set GALLERY_INTEGRATION=1 to run against a live Algolia app (~/.algolia/credentials)")
	}
	indexName := "test-" + uuid.New().String()
	builder := With().Index(indexName)
	client, err := builder.newAdminClient()
	assert.NoError(t, err)
	indexes, err := client.ListIndexes()
	assert.NoError(t, err)
	originalLength := len(indexes)

	object := algoliasearch.Object{
		"testKey":  "value",
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
	assert.Equal(t, 1, res.NbHits)

	client.DeleteIndex(indexName)
}
