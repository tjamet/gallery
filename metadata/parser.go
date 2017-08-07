package metadata

import (
	"os/exec"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
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
	Path string
}

func FromFile(path string) *Metadata {
	return &Metadata{
		Path: path,
	}
}

func (m *Metadata) GetMetadata() (map[string]interface{}, error) {
	stat, err := os.Stat(m.Path)
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, fmt.Errorf("Cannot read metadata of a directory %s", m.Path)
	}
	r := map[string]interface{}{}
	exifList := []map[string]interface{}{}
	cmd := exec.Command("exiftool", "-json", m.Path)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(out.Bytes(), &exifList)
	if err != nil {
		return nil, err
	}
	exif := exifList[0]
	for _, key := range (passThru) {
		value, ok := exif[key]
		if ok {
			r[key] = value
		}
	}

	for old, new := range (mapping) {
		value, ok := exif[old]
		if ok {
			r[new] = value
		}
	}

	return r, nil
}
