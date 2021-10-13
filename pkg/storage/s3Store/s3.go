// Клиент хранилища S3.
// Выполняет контракт storage.Interface.
package s3Store

import (
	"bytes"
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Хранилище S3.
type DB struct {
	client *s3.Client
}

// New создаёт клиент хранилища S3.
func New(ctx context.Context, key, secret, region, url string) (*DB, error) {
	cred := credentials.NewStaticCredentialsProvider(key, secret, "")
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(cred),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, err
	}
	db := DB{
		client: s3.NewFromConfig(cfg),
	}
	return &db, nil
}

func (db *DB) Upload(ctx context.Context, bucket, key, contentType string, content []byte) error {
	reader := bytes.NewReader(content)

	params := &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &key,
		ContentType: &contentType,
		Body:        reader,
	}

	_, err := db.client.PutObject(ctx, params)

	return err
}

// Download читает файл из хранилища.
func (db *DB) Download(ctx context.Context, bucket, key string) ([]byte, error) {
	params := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}

	output, err := db.client.GetObject(ctx, params)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, err
	}

	err = output.Body.Close()
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Delete удаляет файл из хранилища.
func (db *DB) Delete(ctx context.Context, bucket, key string) error {
	params := &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}

	_, err := db.client.DeleteObject(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

// ObjectSize возвращает размер объекта.
func (db *DB) ObjectSize(ctx context.Context, bucket, key string) (int64, error) {
	params := &s3.HeadObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}

	output, err := db.client.HeadObject(ctx, params)
	if err != nil {
		return 0, err
	}

	return output.ContentLength, nil
}
