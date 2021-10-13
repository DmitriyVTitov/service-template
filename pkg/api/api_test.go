package api

import (
	"git.sbercloud.tech/products/sberdisk/sberdisk-b2c/go_service_template/pkg/storage/memdb"
	"os"
	"testing"
)

var api *API

func TestMain(m *testing.M) {
	api = New(memdb.New())
	os.Exit(m.Run())
}
