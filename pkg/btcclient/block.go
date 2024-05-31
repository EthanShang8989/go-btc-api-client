package btcclient

import "fmt"

// GetBlock fetches information about a block by hash
func (c *Client) GetBlockByhash(hash string) (*BlockInfo, error) {
	var block BlockInfo
	err := c.sendRequest(fmt.Sprintf("/block/%s", hash), &block)
	return &block, err
}

// GetBlockHeaderHex fetches the hex-encoded block header by hash
func (c *Client) GetBlockHeaderHex(hash string) (string, error) {
	var header string
	err := c.sendRequest(fmt.Sprintf("/block/%s/header", hash), &header)
	return header, err
}

// GetBlockStatus fetches the status of a block by hash
func (c *Client) GetBlockStatus(hash string) (BlockStatus, error) {
	var status BlockStatus
	err := c.sendRequest(fmt.Sprintf("/block/%s/status", hash), &status)
	return status, err
}

// GetBlockTransactions fetches a list of transactions in a block by hash(up to 25 transactions beginning at start_index).
func (c *Client) GetBlockTransactions(hash string, startIndex ...int) ([]Transaction, error) {
	endpoint := fmt.Sprintf("/block/%s/txs", hash)
	if len(startIndex) > 0 {
		endpoint += fmt.Sprintf("/%d", startIndex[0])
	}
	var txs []Transaction
	err := c.sendRequest(endpoint, &txs)
	return txs, err
}

// GetBlockTxids fetches a list of all txids in a block by hash
func (c *Client) GetBlockTxids(hash string) ([]string, error) {
	var txids []string
	err := c.sendRequest(fmt.Sprintf("/block/%s/txids", hash), &txids)
	return txids, err
}

// GetBlockAllTransactions fetches all transactions in a block by hash. WARM: It may take a long time if there are many tx at one block.
func (c *Client) GetBlockAllTransactions(hash string) ([]Transaction, error) {
	// Get all txids in the block
	txids, err := c.GetBlockTxids(hash)
	if err != nil {
		return nil, err
	}
	// Fetch each transaction by txid
	var transactions []Transaction
	for _, txid := range txids {
		tx, err := c.GetTransaction(txid)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, *tx)
	}

	return transactions, nil
}

// GetBlockTransactionByIndex fetches the txid at the specified index within a block
func (c *Client) GetTxidByBlockhashAndIndex(hash string, index int) (string, error) {
	var tx string
	err := c.sendRequest(fmt.Sprintf("/block/%s/txid/%d", hash, index), &tx)
	return tx, err
}

// GetBlockRaw fetches the raw block representation in binary by hash
func (c *Client) GetBlockRaw(hash string) ([]byte, error) {
	return c.sendRequestForBinary(fmt.Sprintf("/block/%s/raw", hash))
}

// GetBlockHashByHeight fetches the hash of the block currently at the specified height
func (c *Client) GetBlockHashByHeight(height int) (string, error) {
	var hash string
	err := c.sendRequest(fmt.Sprintf("/block-height/%d", height), &hash)
	return hash, err
}

// GetNewestBlocks fetches the 10 newest blocks starting at the tip or at the specified start height
func (c *Client) GetNewestBlocks(startHeight ...int) ([]BlockInfo, error) {
	endpoint := "/blocks"
	if len(startHeight) > 0 {
		endpoint += fmt.Sprintf("/%d", startHeight[0])
	}
	var blocks []BlockInfo
	err := c.sendRequest(endpoint, &blocks)
	return blocks, err
}

// GetTipBlockHeight fetches the height of the last block
func (c *Client) GetTipBlockHeight() (int, error) {
	var height int
	err := c.sendRequest("/blocks/tip/height", &height)
	return height, err
}

// GetTipBlockHash fetches the hash of the last block
func (c *Client) GetTipBlockHash() (string, error) {
	var hash string
	err := c.sendRequest("/blocks/tip/hash", &hash)
	return hash, err
}
