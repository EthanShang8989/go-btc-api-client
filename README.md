# GO BTC API CLIENT

## Overview

This repository is a client library that encapsulates HTTP requests to a third-party API. It provides a convenient interface for interacting with the API, specifically designed for querying and retrieving information from the Bitcoin blockchain. With this library, developers can easily integrate Bitcoin functionality into their applications by making HTTP requests to the API and accessing a wide range of features, including querying blockchain data, retrieving transaction details, checking account balances, and even sending Bitcoin transactions securely.

**Note: This library is currently in an unstable version. Interfaces and types are subject to change as we continue to develop and improve the library.**

## Requirements

- [Go](https://golang.org/) (Version 1.14 or later recommended)

## Installation

```
 go get github.com/EthanShang8989/go-btc-api-client
```

Below is a simple example demonstrating how to use `btcclient` to fetch all transactions in a block:

```
package main

import (
	"fmt"
	"log"
	"github.com/EthanShang8989/go-btc-api-client/btcclient"
)

func main() {
	// Create API client
	client := btcclient.NewClient("https://blockstream.info/testnet/api/")

	// Block hash
	blockHash := "0000000010942ddf9a42bf4b987867badad7c86bce24d28b2bd5cc459ef64c81"

	// Fetch all transactions in the block
	transactions, err := client.GetAllBlockTransactions(blockHash)
	if err != nil {
		log.Fatalf("Failed to get transactions: %v", err)
	}

	// Print transaction information
	for _, tx := range transactions {
		fmt.Printf("Transaction ID: %s\n", tx.Txid)
	}
}
```

## Compatibility

Fully compatible with the Blockstream API and partially compatible with the Mempool API. Note that some Mempool API endpoints are not yet supported.

## Future Work

We are actively working on expanding the compatibility and functionality of the `go-btc-api-client`. The planned features include:

- Full support for all Mempool API endpoints.
- Support for liquid 
- Partial compatibility with btcd types,

Stay tuned for updates and feel free to contribute to these enhancements!

## Contributing

Contributions to the go-btc-api-client are welcome. Please feel free to fork the repository, make changes, and submit pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.