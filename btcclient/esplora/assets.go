package esplora

import (
	"fmt"
	"strconv"
)

// GetAssetInfo fetches information about an asset
func (c *EsploraClient) GetUserIssuedAssetInfo(assetID string) (*UserIssuedAsset, error) {
	var asset UserIssuedAsset
	err := c.sendRequest(fmt.Sprintf("/asset/%s", assetID), &asset)
	return &asset, err
}

// GetAssetInfo fetches information about an asset
func (c *EsploraClient) GetNativeAssetInfo(assetID string) (*NativeAsset, error) {
	var asset NativeAsset
	err := c.sendRequest(fmt.Sprintf("/asset/%s", assetID), &asset)
	return &asset, err
}

// GetAssetTransactions fetches transactions associated with the specified asset
func (c *EsploraClient) GetAssetTransactions(assetID string) ([]Transaction, error) {
	var txs []Transaction
	err := c.sendRequest(fmt.Sprintf("/asset/%s/txs", assetID), &txs)
	return txs, err
}

// GetAssetMempoolTransactions fetches mempool transactions associated with the specified asset
func (c *EsploraClient) GetAssetMempoolTransactions(assetID string) ([]Transaction, error) {
	var txs []Transaction
	err := c.sendRequest(fmt.Sprintf("/asset/%s/txs/mempool", assetID), &txs)
	return txs, err
}

// GetAssetChainTransactions fetches confirmed transactions associated with the specified asset
func (c *EsploraClient) GetAssetChainTransactions(assetID string, lastSeen ...string) ([]Transaction, error) {
	endpoint := fmt.Sprintf("/asset/%s/txs/chain", assetID)
	if len(lastSeen) > 0 {
		endpoint += fmt.Sprintf("/%s", lastSeen[0])
	}
	var txs []Transaction
	err := c.sendRequest(endpoint, &txs)
	return txs, err
}

// GetAssetSupply fetches the current total supply of the specified asset
func (c *EsploraClient) GetAssetSupply(assetID string) (*uint64, error) {
	var supply string
	err := c.sendRequest(fmt.Sprintf("/asset/%s/supply", assetID), &supply)
	res, err := strconv.ParseUint(supply, 10, 64)
	if err != nil {
		return nil, err
	}

	return &res, err
}

// GetAssetSupplyDecimal fetches the current total supply of the specified asset in decimal
func (c *EsploraClient) GetAssetSupplyDecimal(assetID string) (*uint64, error) {
	var Decimal string
	err := c.sendRequest(fmt.Sprintf("/asset/%s/supply/decimal", assetID), &Decimal)
	res, err := strconv.ParseUint(Decimal, 10, 64)
	if err != nil {
		return nil, err
	}
	return &res, err
}

// GetAssetsRegistry fetches the list of issued assets in the asset registry
// start_index: the start index to use for paging. defaults to 0.
// limit: maximum number of assets to return. defaults to 25, maximum 100.
// sort_field: field to sort assets by. one of name, ticker or domain. defaults to ticker.
// sort_dir: sorting direction. one of asc or desc. defaults to asc.
func (c *EsploraClient) GetAssetsRegistry(startIndex int, limit int, sortField string, sortDir string) ([]UserIssuedAsset, error) {
	endpoint := "/assets/registry"
	params := fmt.Sprintf("?start_index=%d&limit=%d&sort_field=%s&sort_dir=%s", startIndex, limit, sortField, sortDir)
	var assets []UserIssuedAsset
	err := c.sendRequest(endpoint+params, &assets)
	return assets, err
}
