package storage

import "context"

// Interface определяет контракт на работу с хранилищем.
type Interface interface {
	// Download читает файл из хранилища.
	Download(ctx context.Context, bucket, key string) ([]byte, error)

	// Upload записывает файл в хранилище.
	Upload(ctx context.Context, bucket, key, contentType string, content []byte) error

	// Delete удаляет файл из хранилища.
	Delete(ctx context.Context, bucket, key string) error

	// ObjectSize возвращает размер объекта.
	ObjectSize(ctx context.Context, bucket, key string) (int64, error)
}
