package mempoolspace_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestGetBlock(t *testing.T) {
	block, err := realClient.GetBlockHeaderByhash("0000000000000002a0324be1eb5c7496a41251557086682f33b78a4440320fa8")
	assert.NoError(t, err)
	spew.Dump(block)
	assert.Equal(t, "0000000000000002a0324be1eb5c7496a41251557086682f33b78a4440320fa8", block.ID)
}

func TestGetBlockHeader(t *testing.T) {
	header, err := realClient.GetBlockHeaderHex("0000000000000002a0324be1eb5c7496a41251557086682f33b78a4440320fa8")
	assert.NoError(t, err)
	assert.NotEmpty(t, header)
	spew.Dump(header)
}

func TestGetBlockStatus(t *testing.T) {
	status, err := realClient.GetBlockStatus("0000000000000002a0324be1eb5c7496a41251557086682f33b78a4440320fa8")
	assert.NoError(t, err)
	assert.NotNil(t, status)
	spew.Dump(status)
}

func TestGetBlockTransactions(t *testing.T) {
	txs, err := realClient.GetBlockTransactions("0000000000000002a0324be1eb5c7496a41251557086682f33b78a4440320fa8")
	assert.NoError(t, err)
	assert.NotEmpty(t, txs)
	spew.Dump(txs)
}

func TestGetBlockTxids(t *testing.T) {
	txids, err := realClient.GetBlockAllTxids("0000000000000002a0324be1eb5c7496a41251557086682f33b78a4440320fa8")
	assert.NoError(t, err)
	assert.NotEmpty(t, txids)
	spew.Dump(txids)
}

func TestGetBlockAllTxs(t *testing.T) {
	txs, err := realClient.GetBlockAllTransactions("0000000010942ddf9a42bf4b987867badad7c86bce24d28b2bd5cc459ef64c81")
	assert.NoError(t, err)
	assert.NotEmpty(t, txs)
	spew.Dump(txs)
}

func TestGetBlockTransactionByIndex(t *testing.T) {
	txid, err := realClient.GetTxidByBlockhashAndIndex("0000000000000002a0324be1eb5c7496a41251557086682f33b78a4440320fa8", 1)
	assert.NoError(t, err)
	assert.NotNil(t, txid)
	spew.Dump(txid)
}

func TestGetBlockRaw(t *testing.T) {
	raw, err := realClient.GetBlockRaw("0000000000000002a0324be1eb5c7496a41251557086682f33b78a4440320fa8")
	assert.NoError(t, err)
	assert.NotEmpty(t, raw)
}

func TestGetBlockHashByHeight(t *testing.T) {
	hash, err := realClient.GetBlockHashByHeight(123456)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	spew.Dump(hash)
}

func TestGetNewestBlocks(t *testing.T) {
	blocks, err := realClient.GetRecentBlocks()
	assert.NoError(t, err)
	assert.NotEmpty(t, blocks)
	spew.Dump(blocks)
}

func TestGetTipBlockHeight(t *testing.T) {
	height, err := realClient.GetTipBlockHeight()
	assert.NoError(t, err)
	assert.True(t, height > 0)
	spew.Dump(height)
}

func TestGetTipBlockHash(t *testing.T) {
	hash, err := realClient.GetTipBlockHash()
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	spew.Dump(hash)
}
