package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubTransactionStore struct {
	transactions map[string]Transaction
}

func (s *StubTransactionStore) GetTransaction(transactionid string) Transaction {
	return s.transactions[transactionid]
}

func TestGetTransactions(t *testing.T) {
	store := &StubTransactionStore{
		transactions: map[string]Transaction{
			"1": Transaction{
				Amount: 100,
			},
			"2": Transaction{
				Amount: 200,
			},
		},
	}

	server := &TransactionServer{
		store: store,
	}

	t.Run("Return transaction 1", func(t *testing.T) {
		request := newGetTransactionRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := parseResponseBody(response.Body)
		want := 100

		assertResponseBody(t, got, want)

	})

	t.Run("Return transaction 2", func(t *testing.T) {
		request := newGetTransactionRequest("2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := parseResponseBody(response.Body)
		want := 200

		assertResponseBody(t, got, want)

	})

	t.Run("Return 404 on missing transaction", func(t *testing.T) {
		request := newGetTransactionRequest("3")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := 404

		assertResponseBody(t, got, want)

	})
}

func newGetTransactionRequest(transactionid string) *http.Request {
	url := fmt.Sprintf("/transactions/%s", transactionid)

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	return req
}

func assertResponseBody(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func parseResponseBody(w io.Reader) int {
	var tran Transaction
	json.NewDecoder(w).Decode(&tran)

	return tran.Amount
}
