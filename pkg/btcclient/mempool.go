package btcclient

// GetMempoolInfo fetches mempool backlog statistics
func (c *Client) GetMempoolInfo() (MempoolInfo, error) {
	var stats MempoolInfo
	err := c.sendRequest("/mempool", &stats)
	return stats, err
}

// GetMempoolTxids fetches the full list of txids in the mempool
func (c *Client) GetMempoolTxids() ([]string, error) {
	var txids []string
	err := c.sendRequest("/mempool/txids", &txids)
	return txids, err
}

// GetRecentMempoolTxs fetches a list of the last 10 transactions to enter the mempool
func (c *Client) GetRecentMempoolTxs() ([]Transaction, error) {
	var txs []Transaction
	err := c.sendRequest("/mempool/recent", &txs)
	return txs, err
}
