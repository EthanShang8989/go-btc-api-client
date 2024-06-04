package esplora

import (
	"fmt"
)

// GetTransaction fetches transaction details by txid
func (c *EsploraClient) GetTransaction(txid string) (*Transaction, error) {
	var tx Transaction
	err := c.sendRequest(fmt.Sprintf("/tx/%s", txid), &tx)
	return &tx, err
}

// GetTransactionStatus fetches the transaction confirmation status
func (c *EsploraClient) GetTransactionStatus(txid string) (map[string]interface{}, error) {
	var status map[string]interface{}
	err := c.sendRequest(fmt.Sprintf("/tx/%s/status", txid), &status)
	return status, err
}

// GetTransactionHex fetches the raw transaction in hex
func (c *EsploraClient) GetTransactionHex(txid string) (string, error) {
	var hex string
	err := c.sendRequest(fmt.Sprintf("/tx/%s/hex", txid), &hex)
	return hex, err
}

// GetTransactionRaw fetches the raw transaction as binary data
func (c *EsploraClient) GetTransactionRaw(txid string) ([]byte, error) {
	raw, err := c.sendRequestForBinary(fmt.Sprintf("/tx/%s/raw", txid))
	return raw, err
}

// GetTransactionMerkleBlockProof fetches a merkle inclusion proof using bitcoind's merkleblock format
func (c *EsploraClient) GetTransactionMerkleBlockProof(txid string) (string, error) {
	proof, err := c.sendRequestForBinary(fmt.Sprintf("/tx/%s/merkleblock-proof", txid))
	return string(proof), err
}

// GetTransactionMerkleProof fetches a merkle inclusion proof using Electrum's format
func (c *EsploraClient) GetTransactionMerkleProof(txid string) (MerkleProof, error) {
	var proof MerkleProof
	err := c.sendRequest(fmt.Sprintf("/tx/%s/merkle-proof", txid), &proof)
	return proof, err
}

// GetTransactionOutputSpendingStatus fetches the spending status of a transaction output
func (c *EsploraClient) GetTransactionOutputSpendingStatus(txid string, vout int) (bool, error) {
	var status map[string]bool
	err := c.sendRequest(fmt.Sprintf("/tx/%s/outspend/%d", txid, vout), &status)
	return status["spent"], err
}

// GetTransactionOutputsSpendingStatus fetches the spending status of all transaction outputs
func (c *EsploraClient) GetTransactionOutputsSpendingStatus(txid string) ([]bool, error) {
	var status []map[string]interface{}
	err := c.sendRequest(fmt.Sprintf("/tx/%s/outspends", txid), &status)
	if err != nil {
		return nil, err
	}

	spentStatuses := make([]bool, len(status))
	for i, s := range status {
		if spent, ok := s["spent"].(bool); ok {
			spentStatuses[i] = spent
		} else {
			return nil, fmt.Errorf("unexpected data format")
		}
	}

	return spentStatuses, nil
}

// BroadcastTransaction broadcasts a raw transaction to the network
func (c *EsploraClient) BroadcastTransaction(rawTx string) (string, error) {
	endpoint := "/tx"
	headers := make(map[string]string)
	headers["Content-Type"] = "text/plain"
	var result string
	err := c.SendPostRequest(endpoint, headers, rawTx, &result)
	if err != nil {
		return "", err
	}
	return result, nil
}
