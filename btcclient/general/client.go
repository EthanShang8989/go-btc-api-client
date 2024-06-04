package general

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

// Client struct represents the client to interact with the API
type Client struct {
	url    string
	client *resty.Client
}

func NewClient(url string) *Client {
	return &Client{
		url:    url,
		client: resty.New(),
	}
}

// SendRequest sends an HTTP request and parses the response
func (c *Client) SendRequest(endpoint string, headers map[string]string, result interface{}) error {
	request := c.client.R()

	// Iterate over the header map and set each key-value pair as a header
	for key, value := range headers {
		request.SetHeader(key, value)
	}

	// Perform the GET request
	resp, err := request.Get(c.url + endpoint)
	if err != nil {
		return err
	}

	// Check the status code of the response
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

// SendRequestForBinary sends an HTTP request and returns the binary response
func (c *Client) SendRequestForBinary(endpoint string, headers map[string]string) ([]byte, error) {
	request := c.client.R()

	// Iterate over the header map and set each key-value pair as a header
	for key, value := range headers {
		request.SetHeader(key, value)
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

// SendPostRequest sends an HTTP POST request and parses the response
func (c *Client) SendPostRequest(endpoint string, headers map[string]string, body interface{}, result interface{}) error {
	request := c.client.R()

	// Iterate over the header map and set each key-value pair as a header
	for key, value := range headers {
		request.SetHeader(key, value)
	}

	// Set the request body as JSON
	request.SetBody(body)

	// Perform the POST request
	resp, err := request.Post(c.url + endpoint)
	if err != nil {
		return err
	}

	// Check the status code of the response
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
