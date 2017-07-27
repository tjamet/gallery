package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type clientBuilder interface {
	SharedConfigEnable() clientBuilder
	Region(region string) clientBuilder
	Context(c *aws.Context) clientBuilder
	Bucket(bucket string) clientBuilder
	UploaderWith() s3Uploader
}

type builder struct {
	config  []*aws.Config
	context *aws.Context
	options session.Options
	bucket  string
	region  string
}

func With() *builder {
	return &builder{
		region: "eu-west-1",
	}
}

func (b *builder) session() (*session.Session, error) {

	b.options.Config.MergeIn(b.config...)
	s, err := session.NewSessionWithOptions(b.options)
	if err != nil {
		return nil, err
	}
	if b.bucket != "" {
		region, err := s3manager.GetBucketRegion(*b.context, s, b.bucket, b.region)
		if err != nil {
			return nil, err
		}
		b.region = region
	}
	return s, nil
}

func (b *builder) SharedConfigEnable() *builder {
	b.options.SharedConfigState = session.SharedConfigEnable
	return b
}

func (b *builder) Region(region string) *builder {
	b.region = region
	return b
}

func (b *builder) Context(c *aws.Context) *builder {
	b.context = c
	return b
}

func (b *builder) Bucket(bucket string) *builder {
	b.bucket = bucket
	return b
}

func (b *builder) buildS3Client() (*s3.S3, error) {
	if b.context == nil {
		c := aws.BackgroundContext()
		b.context = &c
	}
	s, err := b.session()
	if err != nil {
		return nil, err
	}
	return s3.New(s, aws.NewConfig().WithRegion(b.region)), nil
}

func (b *builder) UploaderWith() *uploader {
	cl, err := b.buildS3Client()
	return UploaderWith().
		Error(err).
		Client(cl).Bucket(b.bucket)
}
