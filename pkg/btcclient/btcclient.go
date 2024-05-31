package btcclient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

// Client struct represents the client to interact with the API
type Client struct {
	url    string
	apiKey string
	client *resty.Client
}

// NewClient creates a new instance of Client
//
// example:
//
// client := btcclient.NewClient("https://blockstream.info/testnet/api/")
//
//	blockHash := "0000000010942ddf9a42bf4b987867badad7c86bce24d28b2bd5cc459ef64c81"
//	transactions, err := client.GetBlockAllTransactions(blockHash)
//	if err != nil {
//		log.Fatalf("Failed to get transactions: %v", err)
//	}
//	for _, tx := range transactions {
//		fmt.Printf("Transaction ID: %s\n", tx.Txid)
//	}
//
// TODO:add token if nessesary
func NewClient(url string) *Client {
	var key string
	// if len(apiKey) > 0 {
	// 	key = apiKey[0]
	// }

	return &Client{
		url:    url,
		apiKey: key,
		client: resty.New(),
	}
}

// sendRequest sends an HTTP request and parses the response
func (c *Client) sendRequest(endpoint string, result interface{}) error {
	request := c.client.R()

	if c.apiKey != "" {
		request.SetHeader("Authorization", "Bearer "+c.apiKey)
	}

	resp, err := request.Get(c.url + endpoint)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("error: status code %d", resp.StatusCode())
	}

	// Try to unmarshal JSON response
	if err := json.Unmarshal(resp.Body(), result); err != nil {
		// If JSON unmarshal fails, return the raw response as a string
		if strResult, ok := result.(*string); ok {
			*strResult = string(resp.Body())
			return nil
		}
		return err
	}
	return nil
}

// sendRequestForBinary sends an HTTP request and returns the binary response
func (c *Client) sendRequestForBinary(endpoint string) ([]byte, error) {
	request := c.client.R()

	if c.apiKey != "" {
		request.SetHeader("Authorization", "Bearer "+c.apiKey)
	}

	resp, err := request.Get(c.url + endpoint)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode())
	}

	return resp.Body(), nil
}

// GetBlockHeight fetches the current block height from the API
func (c *Client) GetBlockHeight() (int, error) {
	var height int
	err := c.sendRequest("/blocks/tip/height", &height)
	return height, err
}
