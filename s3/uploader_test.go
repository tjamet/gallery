package s3

import (
	"testing"
	"strings"
	"github.com/stretchr/testify/assert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

func TestUpload(t *testing.T) {
	bucketName := uuid.New().String()
	region := "eu-west-1"
	path := "test"

	builder := With().
		SharedConfigEnable().
		Region(region)

	client, err := builder.buildS3Client()
	assert.NoError(t, err)

	mkBucket := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(region),
		},
	}

	_, err = client.CreateBucket(mkBucket)
	assert.NoError(t, err)

	builder.Bucket(bucketName)
	err = builder.UploaderWith().
		Path(path).
		Body(strings.NewReader("hello world")).
		Upload()
	assert.NoError(t, err)

	rmObj := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
	}
	_, err = client.DeleteObject(rmObj)
	assert.NoError(t, err)

	rmBucket := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}
	_, err = client.DeleteBucket(rmBucket)
	assert.NoError(t, err)
}
