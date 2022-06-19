package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTransactions(t *testing.T) {

	request, _ := http.NewRequest(http.MethodGet, "/transactions/1", nil)
	response := httptest.NewRecorder()

	TransactionServer(response, request)

	got := response.Body.String()
	want := "100"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
