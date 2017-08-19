package s3

import (
	"io"
	"fmt"
)

type DryUploader struct {
	path string
}

type DryBuilder struct{}

func (b *DryBuilder) UploaderWith() s3Uploader {
	return &DryUploader{}
}

func (u *DryUploader) SetBody(body io.ReadSeeker) {}
func (u *DryUploader) SetName(name string)        {}
func (u *DryUploader) Upload() error {
	fmt.Printf("Would upload file to %s\n", u.path)
	return nil
}
