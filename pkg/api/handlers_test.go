package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI_fileInfoHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/files/10/info", nil)
	rr := httptest.NewRecorder()

	api.r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	got := string(body)
	want := "OK"
	if got != want {
		t.Errorf("получен ответ %v, ожидалось %v", got, want)
	}
}
