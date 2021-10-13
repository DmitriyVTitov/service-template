// Исполняемый пакет микросервиса.
package main

import (
	"net/http"
	"os"

	"git.sbercloud.tech/products/sberdisk/sberdisk-b2c/go_service_template/pkg/api"
	"git.sbercloud.tech/products/sberdisk/sberdisk-b2c/go_service_template/pkg/storage"
	"git.sbercloud.tech/products/sberdisk/sberdisk-b2c/go_service_template/pkg/storage/memdb"
)

// Микросервис.
type service struct {
	db   storage.Interface
	api  *api.API
	conf *config
}

// Конфигурация микросервиса.
type config struct {
	dbConnString string
}

func main() {
	svc := buildService()
	http.ListenAndServe(":80", svc.api.Router())
}

// Конструктор микросервиса.
func buildService() *service {
	db := memdb.New()
	api := api.New(db)
	s := service{
		db:   db,
		api:  api,
		conf: new(config),
	}
	s.parseConfig()
	return &s
}

// Заполнение конфигурации микросервиса данными
// из переменных окружения.
func (s *service) parseConfig() {
	s.conf.dbConnString = os.Getenv("dbconnstr")
}
