// Реализация контракта storage.Interface в памяти для тестов.
package memdb

import (
	"context"
)

type MemDB struct{}

func New() *MemDB {
	return new(MemDB)
}

// Download читает файл из хранилища.
func (*MemDB)Download(ctx context.Context, bucket, key string) ([]byte, error) {
	return nil, nil
}

// Upload записывает файл в хранилище.
func (*MemDB)Upload(ctx context.Context, bucket, key, contentType string, content []byte) error {
	return nil
}

// Delete удаляет файл из хранилища.
func (*MemDB)Delete(ctx context.Context, bucket, key string) error {
	return nil
}

// ObjectSize возвращает размер объекта.
func (*MemDB)ObjectSize(ctx context.Context, bucket, key string) (int64, error) {
	return 0,nil
}