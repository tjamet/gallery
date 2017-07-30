package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"io"
	"net/http"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"fmt"
	"crypto/sha256"
)

type s3Uploader interface {
	SetBody(body io.ReadSeeker)
	Upload() error
}

func wrap(format string, err error, args ...interface{}) error {
	if err!= nil {
		format = format + ": %s"
		args = append(args, err.Error())
		return fmt.Errorf(format, args...)
	}
	return nil
}

type uploader struct {
	bucket        *string
	key           *string
	body          io.ReadSeeker
	contentLength *int64
	contentType   *string
	client        *s3.S3
	err           []error
}

func (u *uploader) Client(client *s3.S3) *uploader {
	u.client = client
	return u
}

func (u *uploader) Bucket(bucket string) *uploader {
	u.bucket = aws.String(bucket)
	return u
}

func (u *uploader) Key(key string) *uploader {
	u.key = aws.String(key)
	return u
}

func (u *uploader) Path(key string) *uploader {
	return u.Key(key)
}

func (u *uploader) Error(err error) *uploader {
	if err != nil {
		u.err = append(u.err, err)
	}
	return u
}

func (u *uploader) Body(body io.ReadSeeker) *uploader {
	u.body = body
	data, err := ioutil.ReadAll(body)
	u.Error(wrap("Failed to read content", err))
	_, err = body.Seek(0, io.SeekStart)
	u.Error(wrap("Failed to read content", err))
	u.contentLength = aws.Int64(int64(len(data)))
	u.ContentType(http.DetectContentType(data))
	if u.key == nil {
		h := sha256.New()
		h.Write(data)
		h.Sum(nil)
		u.key = aws.String(fmt.Sprintf("%x", h.Sum(nil)))
	}
	return u
}

func (u *uploader) ContentType(contentType string) *uploader {
	u.contentType = aws.String(contentType)
	return u
}

func UploaderWith() *uploader {
	return &uploader{}
}

func (u *uploader) SetBody(body io.ReadSeeker) {
	u.Body(body)
}

func (u *uploader) Upload() error {
	if len(u.err) > 0 {
		// TODO: format all errors
		return u.err[0]
	}
	params := &s3.PutObjectInput{
		Bucket:        u.bucket,
		Key:           u.key,
		Body:          u.body,
		ContentLength: u.contentLength,
		ContentType:   u.contentType,
	}
	_, err := u.client.PutObject(params)
	return err
}

func (u *uploader) GetKey() string {
	return *u.key
}

func (u *uploader) GetName() string {
	return u.GetKey()
}