// REST API микросервиса.
package api

import (
	"net/http"
	"net/http/pprof"

	"git.sbercloud.tech/products/sberdisk/sberdisk-b2c/go_service_template/pkg/storage"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// API содержит зависимости.
type API struct {
	db storage.Interface
	r  *mux.Router
}

// Метрики для системы Prometheus.
var (
	requestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "requests_total",
		Help: "Общее количество запросов к API.",
	})
)

// New создаёт объект API.
func New(db storage.Interface) *API {
	api := API{
		db: db,
		r:  mux.NewRouter(),
	}
	api.middleware()
	api.endpoints()
	return &api
}

// Регистрация промежуточного ПО.
func (api *API) middleware() {
	api.r.Use(api.headersMiddleware)
	api.r.Use(api.prometheusMiddleware)
}

// Регистрация конечных точек REST API.
func (api *API) endpoints() {
	// обработчик для снятия метрик Prometheus
	api.r.Handle("/metrics", promhttp.Handler())
	// Обработчики для pprof
	api.r.HandleFunc("/debug/pprof/", pprof.Index)
	api.r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	api.r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	api.r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	api.r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	api.r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	api.r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	api.r.Handle("/debug/pprof/block", pprof.Handler("block"))
	api.r.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))

	// API службы
	api.r.HandleFunc("/api/v1/files/{id}/info", api.fileInfoHandler).Methods(http.MethodGet)
}

// Router возвращает маршрутизатор HTTP.
func (api *API) Router() *mux.Router {
	return api.r
}
