package util

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"io"
)

func NewS3Client(cfg aws.Config) *s3.Client {
	return s3.NewFromConfig(cfg)
}

// ListBuckets returns a list of all buckets owned by the authenticated sender of the request.
func ListBuckets(client *s3.Client, ctx context.Context) (*s3.ListBucketsOutput, error) {
	return client.ListBuckets(ctx, &s3.ListBucketsInput{})
}

// ListObjects returns some or all (up to 1000) of the objects in a bucket.
func ListObjects(client *s3.Client, ctx context.Context, bucket, prefix string) (*s3.ListObjectsV2Output, error) {
	return client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String("/"),
	})
}

// PutObject uploads a new object to the specified Amazon S3 bucket.
func PutObject(client *s3.Client, ctx context.Context, bucket, filename string, buf io.Reader) (*s3.PutObjectOutput, error) {
	return client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   buf,
	})
}

// GetObject returns the contents of the specified object.
func GetObject(client *s3.Client, ctx context.Context, bucket, key string) (*s3.GetObjectOutput, error) {
	return client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
}

// DeleteObject removes the null version (if there is one) of an object and inserts a delete marker, which becomes the latest version of the object.
func DeleteObject(client *s3.Client, ctx context.Context, bucket string, key ...string) (*s3.DeleteObjectsOutput, error) {
	var input []types.ObjectIdentifier
	for _, v := range key {
		input = append(input, types.ObjectIdentifier{
			Key: aws.String(v),
		})
	}
	return client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
		Bucket: aws.String(bucket),
		Delete: &types.Delete{
			Objects: input,
			Quiet:   aws.Bool(false),
		},
	})
}
