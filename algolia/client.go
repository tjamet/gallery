package algolia

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/tjamet/gallery/config"
)

type AlgoliaClient interface {
	UpdateObjects(objects ...algoliasearch.Object) error
}

type client struct {
	indexName string
	client    algoliasearch.Client
	indexes   []algoliasearch.Index
	cfg       config.Getter
}

func With() *client {
	g := config.With().
		Path("~/.algolia/credentials").
		Section("default").
		Build()
	return (&client{}).Config(g)
}

func (c *client) Config(cfg config.Getter) *client {
	c.cfg = cfg
	return c
}

func (c *client) Index(name string) *client {
	c.indexName = name
	return c
}

func (c *client) newAdminClient() (algoliasearch.Client, error) {
	appID, err := c.cfg.Get("appID")
	if err != nil {
		return nil, err
	}
	adminSecret, err := c.cfg.Get("adminKey")
	if err != nil {
		return nil, err
	}
	return algoliasearch.NewClient(appID, adminSecret), nil
}

func (c *client) UpdateObjects(object ...algoliasearch.Object) error {
	cl, err := c.newAdminClient()
	if err != nil {
		return err
	}
	index := cl.InitIndex(c.indexName)
	_, err = index.AddObjects(object)
	return err
}

func NewObjectFromMetadata(in map[string]interface{}) algoliasearch.Object {
	r := algoliasearch.Object(in)
	_, ok := r["objectID"]
	if !ok {
		_, ok = r["ID"]
		if ok {
			r["objectID"] = r["ID"]
			delete(r, "ID")
		}
	}
	return r
}
