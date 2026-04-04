package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.etherscan.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connnection")
	_ = client
}
