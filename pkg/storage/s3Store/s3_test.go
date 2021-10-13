package s3Store

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

const (
	file   = "./s3_test.go"
	key    = "s3_test.go"
	bucket = "dtitov-test-bucket"
)

var db *DB

func TestMain(m *testing.M) {
	var err error
	db, err = New(context.Background(),
		os.Getenv("S3_KEY"),
		os.Getenv("S3_SECRET"),
		os.Getenv("S3_REGION"),
		"http://localhost",
	)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func TestDB_Upload(t *testing.T) {
	b, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		ctx     context.Context
		bucket  string
		key     string
		content []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Отправка данного файла в хранилище",
			args: args{
				ctx:     context.Background(),
				bucket:  bucket,
				key:     key,
				content: b,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := db.Upload(tt.args.ctx, tt.args.bucket, tt.args.key, "audio/mpeg",
				tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Upload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_Download(t *testing.T) {
	type args struct {
		ctx    context.Context
		bucket string
		key    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Получение отправленного файла из хранилища",
			args: args{
				ctx:    context.Background(),
				bucket: bucket,
				key:    key,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.Download(tt.args.ctx, tt.args.bucket, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("Download() len == 0")
			}
		})
	}
}

func TestDB_ObjectSize(t *testing.T) {
	type args struct {
		ctx    context.Context
		bucket string
		key    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Получение размера файла",
			args: args{
				ctx:    context.Background(),
				bucket: bucket,
				key:    key,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := ioutil.ReadFile(file)
			if err != nil {
				t.Fatal(err)
			}

			if err := db.Upload(tt.args.ctx, tt.args.bucket, "video/mp4", tt.args.key, b); (err != nil) != tt.wantErr {
				t.Errorf("Upload() error = %v, wantErr %v", err, tt.wantErr)
			}

			objSize, err := db.ObjectSize(tt.args.ctx, tt.args.bucket, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetObjectSize() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := db.Delete(tt.args.ctx, tt.args.bucket, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			if int(objSize) != len(b) {
				t.Errorf("GetObjectSize() Размеры файлов разные")
			}
		})
	}
}

func TestDB_Delete(t *testing.T) {
	type args struct {
		ctx    context.Context
		bucket string
		key    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Удаление файла из хранилища",
			args: args{
				ctx:    context.Background(),
				bucket: bucket,
				key:    key,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := db.Delete(tt.args.ctx, tt.args.bucket, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			_, err := db.Download(context.Background(), bucket, key)
			if err == nil {
				t.Errorf("Delete() Файл существует в хранилище")
			}
		})
	}
}
