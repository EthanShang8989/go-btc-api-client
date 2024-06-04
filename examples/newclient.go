package main

import (
	"fmt"
	"log"

	"github.com/EthanShang8989/go-btc-api-client/btcclient/esplora"
)

func main() {
	client := esplora.NewClient("https://blockstream.info/testnet/api")
	blockHash := "0000000010942ddf9a42bf4b987867badad7c86bce24d28b2bd5cc459ef64c81"
	transactions, err := client.GetBlockAllTransactions(blockHash)
	if err != nil {
		log.Fatalf("Failed to get transactions: %v", err)
	}
	for _, tx := range transactions {
		fmt.Printf("Transaction ID: %s\n", tx.Txid)
	}
}
