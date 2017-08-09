package galery

import (
	"fmt"

	"log"
	"os"

	"github.com/tjamet/gallery/algolia"
	"github.com/tjamet/gallery/metadata"
	"github.com/tjamet/gallery/s3"
)

type Pipe struct {
	UploaderBuilder s3.ClientBuilder
	SearchClient    algolia.AlgoliaClient
	BuildURLs       func(string) map[string]string
}

type entry struct {
	metadata *metadata.Metadata
	path     string
	err      error
}

func (e *entry) HasError() bool {
	return e.err != nil
}

func (p *Pipe) Entry(metadata *metadata.Metadata) *entry {
	return &entry{
		metadata: metadata,
		err:      nil,
	}
}

func (p *Pipe) Upload(e *entry) *entry {
	fmt.Println(e.metadata.Path, e.metadata.ID)
	e.path = e.metadata.ID
	fileUploader := p.UploaderBuilder.UploaderWith()
	fileUploader.SetName(e.path)
	f, err := os.Open(e.metadata.Path)
	if err != nil {
		e.err = err
		return e
	}
	defer f.Close()
	fileUploader.SetBody(f)
	err = fileUploader.Upload()
	if err != nil {
		e.err = err
		return e
	}
	log.Printf("Uploaded %s -> %s", e.metadata.Path, e.path)
	return e
}

func (p *Pipe) Index(e *entry) *entry {
	metadata, err := e.metadata.GetMetadata()
	if err != nil {
		e.err = err
		return e
	}
	object := algolia.NewObjectFromMetadata(metadata)
	object["sizes"] = p.BuildURLs(e.path)
	e.err = p.SearchClient.UpdateObjects(object)
	return e
}
