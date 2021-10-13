package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Обработчики конечных точек REST API.

// fileInfoHandler - пример обработчика.
func (api *API) fileInfoHandler(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["id"]
	id, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = id
	w.Write([]byte("OK"))
}
