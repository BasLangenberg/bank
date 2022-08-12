package main

import (
	"testing"
)

func TestCreateAccount(t *testing.T) {
	t.Run("Create account for J. Doe", func(t *testing.T) {
		account := createAccount("Jane", "Doe")
		if account.firstName != "Jane" {
			t.Errorf("Want: %s, Got %s", "Jane", account.firstName)
		}
	})
}
