package main

import (
	"encoding/json"
	"fmt"
)

type NoEmpty struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"creationyear,omitempty"`
}

type Password struct {
	Name    string `json:"username"`
	Surname string `json:"surname,omitempty"`
	Year    int    `json:"creationyear,omitempty"`
	Pass    string `json:"-"`
}

func main8() {
	noempty := NoEmpty{Name: "Mihalis"}
	password := Password{Name: "Mihalis", Pass: "myPassword"}

	noEmptyVar, err := json.Marshal(&noempty)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("noEmptyVar decoded with value %s\n", noEmptyVar)
	}

	passwordVar, err := json.Marshal(&password)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("password decoded with value %s\n", passwordVar)
	}
}
