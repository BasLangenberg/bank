package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(TransactionServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
