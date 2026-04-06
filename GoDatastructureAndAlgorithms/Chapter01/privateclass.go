package main

import (
	"encoding/json"
	"fmt"
)

type AccountDetails struct {
	id          string
	accountType string
}

type Account1 struct {
	details      *AccountDetails
	CustomerName string
}

func (account *Account1) setDetails(id string, accountType string) {
	account.details = &AccountDetails{id, accountType}
}

func (account *Account1) getId() string {
	return account.details.id
}

func (account *Account1) getAccountType() string {
	return account.details.accountType
}

func main10() {
	var account *Account1 = &Account1{CustomerName: "John Smith"}
	account.setDetails("4532", "current")

	jsonAccount, _ := json.Marshal(account)
	fmt.Println("Private Class hidden", string(jsonAccount))
	fmt.Println("Account id", account.getId())
	fmt.Println("Account Type", account.getAccountType())
}
