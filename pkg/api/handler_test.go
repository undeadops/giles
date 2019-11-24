package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/undeadops/giles/pkg/storage/db/mock"
)

func TestTickerWatchList(t *testing.T) {
	req, err := http.NewRequest("GET", "/tickerwatch/", nil)
	if err != nil {
		t.Fatal(err)
	}

	mDB := &mockdb.MockDB{}

	s := Server{DB: mDB}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.TickerWatcherList)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v expected %v",
			status, http.StatusOK)
	}

	expected := `{"alive": true }`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
