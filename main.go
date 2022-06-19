package main

import (
	"log"
	"net/http"
)

type InMemoryTransactionStore struct{}

func (i *InMemoryTransactionStore) GetTransaction(id string) Transaction {
	return Transaction{
		Amount: 1000,
	}
}

func main() {
	handler := &TransactionServer{&InMemoryTransactionStore{}}
	log.Fatal(http.ListenAndServe(":5000", handler))
}
