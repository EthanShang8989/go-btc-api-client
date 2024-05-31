package btcclient

// GetFeeEstimates fetches fee estimates for different confirmation targets
func (c *Client) GetFeeEstimates() (map[string]float64, error) {
	var estimates map[string]float64
	err := c.sendRequest("/fee-estimates", &estimates)
	return estimates, err
}

//TODO:Compatible with mempool's recommendFee interface
