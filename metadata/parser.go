package metadata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

var passThru []string = []string{
	"Artist",
	"Title",
	"Copyright",
	"Creator",
	"Keywords",
}

var mapping map[string]string = map[string]string{
	"DateTimeOriginal":   "Date",
	"OriginalDocumentID": "ID",
}

type Metadata struct {
	Path     string
	ID       string
	Metadata map[string]interface{}
	err      error
}

func FromFile(path string) *Metadata {
	return &Metadata{
		Path: path,
		err:  fmt.Errorf("Metadata has not yet been loaded, please call the Load function"),
	}
}

func (m *Metadata) HasError() bool {
	return m.err != nil
}

func (m *Metadata) Load() *Metadata {
	stat, err := os.Stat(m.Path)
	if err != nil {
		m.err = err
		return m
	}
	if stat.IsDir() {
		m.err = fmt.Errorf("Cannot read metadata of a directory %s", m.Path)
		return m
	}
	r := map[string]interface{}{}
	exifList := []map[string]interface{}{}
	cmd := exec.Command("exiftool", "-json", m.Path)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Start()
	if err != nil {
		m.err = err
		return m
	}
	err = cmd.Wait()
	if err != nil {
		m.err = err
		return m
	}
	err = json.Unmarshal(out.Bytes(), &exifList)
	if err != nil {
		m.err = err
		return m
	}
	exif := exifList[0]
	for _, key := range passThru {
		value, ok := exif[key]
		if ok {
			r[key] = value
		}
	}

	for old, new := range mapping {
		value, ok := exif[old]
		if ok {
			r[new] = value
		}
	}
	v, ok := r["ID"]
	if ok {
		m.ID = v.(string)
	}

	m.Metadata = r
	m.err = nil
	return m
}

func (m *Metadata) GetMetadata() (map[string]interface{}, error) {
	return m.Metadata, m.err
}

func Load(m *Metadata) *Metadata {
	return m.Load()
}
