package mempoolspace

import (
	"fmt"

	"github.com/EthanShang8989/go-btc-api-client/btcclient/esplora"
)

type MempoolSpaceClient struct {
	*esplora.EsploraClient
}

func NewClient(url string) *MempoolSpaceClient {
	c := esplora.NewClient(url)
	return &MempoolSpaceClient{c}
}

//blockchain API Key (ApiKeyAuth)  API 密钥 (ApiKeyAuth)
// Parameter Name: X-API-Token, in: header.
// 参数名称：X-API-Token，位于：标头中。

// blockstream
// sendRequest sends an HTTP request and parses the response
func (c *MempoolSpaceClient) sendRequest(endpoint string, result interface{}) error {
	emptyMap := make(map[string]string)
	c.Client.SendRequest(endpoint, emptyMap, result)
	return nil
}

// sendRequestForBinary sends an HTTP request and returns the binary response
func (c *MempoolSpaceClient) sendRequestForBinary(endpoint string) ([]byte, error) {
	emptyMap := make(map[string]string)
	return c.Client.SendRequestForBinary(endpoint, emptyMap)
}

func (c *MempoolSpaceClient) GetDifficultyAdjustment() (DifficultyAdjustment, error) {
	var difficultyAdjustment DifficultyAdjustment
	err := c.sendRequest("/v1/difficulty-adjustment", &difficultyAdjustment)
	return difficultyAdjustment, err
}

func (c *MempoolSpaceClient) GetHistoricalPrice(currency string, timestamp int64) (HistoricalPrice, error) {
	url := fmt.Sprintf("/v1/historical-price?currency=%s&timestamp=%d", currency, timestamp)
	var historicalPrice HistoricalPrice
	err := c.sendRequest(url, &historicalPrice)
	return historicalPrice, err
}

func (c *MempoolSpaceClient) GetBtcPrice() (Price, error) {
	var prices Price
	err := c.sendRequest("/v1/prices", &prices)
	return prices, err
}
