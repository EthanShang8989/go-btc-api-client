package esplora_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

var testAssetID = "5550cc4ca28648726fbab819188146f1cc305ac3d0ad64209266f84d78d05be8"

func TestGetUserIssuedAssetInfo(t *testing.T) {
	asset, err := liquidClient.GetUserIssuedAssetInfo(testAssetID)
	assert.NoError(t, err)
	assert.NotNil(t, asset)
	spew.Dump(asset)
}
func TestGetNativeAssetInfo(t *testing.T) {
	asset, err := liquidClient.GetNativeAssetInfo("6f0279e9ed041c3d710a9f57d0c02928416460c4b722ae3457a11eec381c526d")
	assert.NoError(t, err)
	assert.NotNil(t, asset)
	spew.Dump(asset)
}

func TestGetAssetTransactions(t *testing.T) {
	transactions, err := liquidClient.GetAssetTransactions(testAssetID)
	assert.NoError(t, err)
	assert.NotNil(t, transactions)
	assert.True(t, len(transactions) > 0)
	spew.Dump(transactions)
}

func TestGetAssetMempoolTransactions(t *testing.T) {
	transactions, err := liquidClient.GetAssetMempoolTransactions(testAssetID)
	assert.NoError(t, err)
	assert.NotNil(t, transactions)
	spew.Dump(transactions)
}

func TestGetAssetChainTransactions(t *testing.T) {
	transactions, err := liquidClient.GetAssetChainTransactions(testAssetID)
	assert.NoError(t, err)
	assert.NotNil(t, transactions)
	assert.True(t, len(transactions) > 0)
	spew.Dump(transactions)
}

func TestGetAssetSupply(t *testing.T) {
	supply, err := liquidClient.GetAssetSupply(testAssetID)
	assert.NoError(t, err)
	assert.NotNil(t, supply)
	spew.Dump(supply)
}

func TestGetAssetSupplyDecimal(t *testing.T) {
	Decimal, err := liquidClient.GetAssetSupplyDecimal(testAssetID)
	assert.NoError(t, err)
	assert.NotNil(t, Decimal)
	spew.Dump(Decimal)

}

func TestGetAssetsRegistry(t *testing.T) {
	assets, err := liquidClient.GetAssetsRegistry(0, 25, "name", "asc")
	assert.NoError(t, err)
	assert.NotNil(t, assets)
	assert.True(t, len(assets) > 0)
	spew.Dump(assets)
}
