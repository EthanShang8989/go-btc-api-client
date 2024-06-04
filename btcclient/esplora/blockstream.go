package esplora

import (
	"github.com/EthanShang8989/go-btc-api-client/btcclient/general"
)

type EsploraClient struct {
	*general.Client
}

func NewClient(url string) *EsploraClient {
	c := general.NewClient(url)
	return &EsploraClient{c}
}

//blockchain API Key (ApiKeyAuth)  API 密钥 (ApiKeyAuth)
// Parameter Name: X-API-Token, in: header.
// 参数名称：X-API-Token，位于：标头中。

// blockstream
// sendRequest sends an HTTP request and parses the response
func (c *EsploraClient) sendRequest(endpoint string, result interface{}) error {
	emptyMap := make(map[string]string)
	c.Client.SendRequest(endpoint, emptyMap, result)
	return nil
}

// sendRequestForBinary sends an HTTP request and returns the binary response
func (c *EsploraClient) sendRequestForBinary(endpoint string) ([]byte, error) {
	emptyMap := make(map[string]string)
	return c.Client.SendRequestForBinary(endpoint, emptyMap)
}
