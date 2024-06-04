package mempoolspace_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAddressInfo(t *testing.T) {
	info, err := realClient.GetAddressInfo("tb1pau685khfkvselml6l5sqx2la7gfdal34szj37tyvm8rqp58244nqyszcm9")
	assert.NoError(t, err)
	assert.NotNil(t, info)
	require.True(t, info.ChainStats.FundedTxoCount > 1)
	// spew.Dump(info)
}
func TestGetAddressValidation(t *testing.T) {
	info, err := realClient.GetAddressValidation("tb1pau685khfkvselml6l5sqx2la7gfdal34szj37tyvm8rqp58244nqyszcm9")
	assert.NoError(t, err)
	assert.NotNil(t, info)
	require.True(t, info.IsWitness == true)
	require.True(t, info.IsScript == true)
	spew.Dump(info)
}

func TestGetScripthashInfo(t *testing.T) {
	info, err := realClient.GetScripthashInfo("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	assert.NoError(t, err)
	assert.NotNil(t, info)
	spew.Dump(info)
}

func TestGetAddressTxs(t *testing.T) {
	txs, err := realClient.GetAddressTxs("tb1pau685khfkvselml6l5sqx2la7gfdal34szj37tyvm8rqp58244nqyszcm9")
	assert.NoError(t, err)
	assert.NotEmpty(t, txs)
	// spew.Dump(txs)
}

func TestGetScripthashTxs(t *testing.T) {
	txs, err := realClient.GetScripthashTxs("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	assert.NoError(t, err)
	assert.NotEmpty(t, txs)
	spew.Dump(txs)
}

func TestGetAddressNewestTxsByTxid(t *testing.T) {
	txs, err := realClient.GetAddressNewestTxsStartFromTxid("tb1pau685khfkvselml6l5sqx2la7gfdal34szj37tyvm8rqp58244nqyszcm9")
	assert.NoError(t, err)
	assert.NotEmpty(t, txs)
	spew.Dump(txs)
}

func TestGetAllAddressTxs(t *testing.T) {
	txs, err := realClient.GetAllAddressTxs("tb1pau685khfkvselml6l5sqx2la7gfdal34szj37tyvm8rqp58244nqyszcm9")
	assert.NoError(t, err)
	assert.NotEmpty(t, txs)
	spew.Dump(txs)
}

func TestGetScripthashNewestTxsStartFromTxid(t *testing.T) {
	txs, err := realClient.GetScripthashNewestTxsStartFromTxid("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	assert.NoError(t, err)
	assert.NotEmpty(t, txs)
	spew.Dump(txs)
}

//TODO: it take too long time to execute,replace scripthash
// func TestGetAllScripthashTxs(t *testing.T) {
// 	txs, err := realClient.GetAllScripthashTxs("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, txs)
// 	spew.Dump(txs)
// }

func TestGetAddressTxsMempool(t *testing.T) {
	txs, err := realClient.GetAddressTxsMempool("tb1pau685khfkvselml6l5sqx2la7gfdal34szj37tyvm8rqp58244nqyszcm9")
	assert.NoError(t, err)
	spew.Dump(txs)
}

func TestGetScripthashTxsMempool(t *testing.T) {
	txs, err := realClient.GetScripthashTxsMempool("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	assert.NoError(t, err)
	spew.Dump(txs)
}

func TestGetAddressUtxos(t *testing.T) {
	utxos, err := realClient.GetAddressUtxos("tb1pau685khfkvselml6l5sqx2la7gfdal34szj37tyvm8rqp58244nqyszcm9")
	assert.NoError(t, err)
	assert.NotEmpty(t, utxos)
	spew.Dump(utxos)
}

func TestGetScripthashUtxos(t *testing.T) {
	utxos, err := realClient.GetScripthashUtxos("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	assert.NoError(t, err)
	assert.NotEmpty(t, utxos)
}

func TestSearchAddressesByPrefix(t *testing.T) {
	addresses, err := realClient.SearchAddressesByPrefix("tb1pau685khfkv")
	assert.NoError(t, err)
	assert.NotEmpty(t, addresses)
	spew.Dump(addresses)
}
