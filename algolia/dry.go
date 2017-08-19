package algolia

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"fmt"
	"encoding/json"
	"bytes"
)

type DryRun struct {}

func (c *DryRun) UpdateObjects(objects ...algoliasearch.Object) error {
	b := bytes.Buffer{}
	enc := json.NewEncoder(&b)
	enc.SetIndent("", "    ")
	_ = enc.Encode(objects)
	fmt.Printf("Would update indices %s\n", b.String())
	return nil
}