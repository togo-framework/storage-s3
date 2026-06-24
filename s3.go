// Package s3store is an AWS S3 driver for togo storage. It implements
// togo.Storage and overrides the default filesystem storage when installed.
// Blank-import + set S3_BUCKET (and the standard AWS_* credentials/region).
//
//	togo install togo-framework/storage-s3
//
// Env: S3_BUCKET (required), AWS_REGION (default us-east-1), AWS_ACCESS_KEY_ID,
// AWS_SECRET_ACCESS_KEY (or any default AWS credential source).
package s3store

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("storage-s3", togo.PriorityService+10, func(k *togo.Kernel) error {
		bucket := os.Getenv("S3_BUCKET")
		if bucket == "" {
			if k.Log != nil {
				k.Log.Warn("storage-s3: S3_BUCKET not set; skipping (using default storage)")
			}
			return nil
		}
		region := envOr("AWS_REGION", "us-east-1")
		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
		if err != nil {
			return fmt.Errorf("storage-s3: load aws config: %w", err)
		}
		k.Storage = &store{cl: s3.NewFromConfig(cfg), bucket: bucket, region: region}
		return nil
	})
}

type store struct {
	cl     *s3.Client
	bucket string
	region string
}

func (s *store) Put(path string, data []byte) error {
	_, err := s.cl.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key(path)),
		Body:   bytes.NewReader(data),
	})
	return err
}

func (s *store) Get(path string) ([]byte, error) {
	out, err := s.cl.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key(path)),
	})
	if err != nil {
		return nil, err
	}
	defer out.Body.Close()
	return io.ReadAll(out.Body)
}

func (s *store) Delete(path string) error {
	_, err := s.cl.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key(path)),
	})
	return err
}

func (s *store) Path(path string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, key(path))
}

func key(p string) string { return strings.TrimPrefix(p, "/") }

func envOr(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
