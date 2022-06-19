package main

import (
	"fmt"
	"net/http"
)

func TransactionServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "100")
}
