package main

import (
	"encoding/json"
	"fmt"
)

type UseAll struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"created"`
}

func main6() {
	useAll := UseAll{
		Name:    "Mike",
		Surname: "Tsoukalos",
		Year:    2020,
	}
	t, err := json.Marshal(&useAll)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value %s\n", t)
	}
	str := `{"username": "M.", "surname": "Ts", "created":2024}`
	jsonRecord := []byte(str)
	temp := UseAll{}
	err = json.Unmarshal(jsonRecord, &temp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Data type: %T with value %v\n", temp, temp)
	}
}
