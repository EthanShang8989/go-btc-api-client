package tests

import "github.com/EthanShang8989/go-btc-api-client/btcclient"

const blockstreamURL = "https://blockstream.info/testnet/api"
const mempoolURL = "https://mempool.space/testnet/api/"

var realClient = btcclient.NewClient(blockstreamURL)
