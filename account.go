package main

type Account struct {
	firstName string
	lastName  string
}

func createAccount(firstname, lastname string) *Account {
	return &Account{
		firstName: firstname,
		lastName:  lastname,
	}
}
