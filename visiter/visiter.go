package visiter

import (
	"os"
	"github.com/chrislusf/glow/flow"
	"log"
	"path/filepath"
)

type VisitInfo struct {
	Path  string
	finfo os.FileInfo
	err   error
}

func Visited(path string, finfo os.FileInfo, err error) *VisitInfo {
	return &VisitInfo{
		Path:  path,
		finfo: finfo,
		err:   err,
	}
}

func (v *VisitInfo) IsDir() bool {
	return v.finfo.IsDir()
}

func New(path string, shard int) *flow.Dataset {
	fc := flow.New()
	fn := func(out chan *VisitInfo) {
		err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
			out <- Visited(path, f, err)
			return nil
		})
		if err != nil {
			log.Printf("Listing folder %s: %v", path, err)
		}
	}
	return fc.Source(fn, shard)
}
