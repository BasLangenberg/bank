package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Transaction struct {
	Amount int
}

type TransactionStore interface {
	GetTransaction(id string) Transaction
}

type TransactionServer struct {
	store TransactionStore
}

func (t *TransactionServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	transaction := strings.TrimPrefix(r.URL.Path, "/transactions/")
	json.NewEncoder(w).Encode(t.store.GetTransaction(transaction))
}
