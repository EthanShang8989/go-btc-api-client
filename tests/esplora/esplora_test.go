package esplora_test

import (
	"github.com/EthanShang8989/go-btc-api-client/btcclient"
	"github.com/EthanShang8989/go-btc-api-client/btcclient/esplora"
)

var realClient = esplora.NewClient(btcclient.MenpoolTestnetURL)
