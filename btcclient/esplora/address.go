package esplora

import (
	"fmt"
)

// GetAddressInfo fetches information about an address
func (c *EsploraClient) GetAddressInfo(address string) (AddressInfo, error) {
	var info AddressInfo
	err := c.sendRequest(fmt.Sprintf("/address/%s", address), &info)
	return info, err
}

// GetScripthashInfo fetches information about a scripthash
func (c *EsploraClient) GetScripthashInfo(hash string) (ScripthashInfo, error) {
	var info ScripthashInfo
	err := c.sendRequest(fmt.Sprintf("/scripthash/%s", hash), &info)
	return info, err
}

// GetAddressTxs fetches transaction history for the specified address
// Get transaction history for the specified address/scripthash, sorted with newest first.
// Returns up to 50 mempool transactions plus the first 25 confirmed transactions.
// You can request more confirmed transactions using GetAllAddressTxs
func (c *EsploraClient) GetAddressTxs(address string) ([]Transaction, error) {
	var txs []Transaction
	err := c.sendRequest(fmt.Sprintf("/address/%s/txs", address), &txs)
	return txs, err
}

// GetScripthashTxs fetches transaction history for the specified scripthash
func (c *EsploraClient) GetScripthashTxs(hash string) ([]map[string]interface{}, error) {
	var txs []map[string]interface{}
	err := c.sendRequest(fmt.Sprintf("/scripthash/%s/txs", hash), &txs)
	return txs, err
}

// GetAddressTxsChain Get confirmed transaction history for the specified address, sorted with newest first.
// Returns 25 transactions per page. More can be requested by specifying the last txid seen by the previous query.
func (c *EsploraClient) GetAddressNewestTxsStartFromTxid(address string, lastSeenTxid ...string) ([]Transaction, error) {
	endpoint := fmt.Sprintf("/address/%s/txs/chain", address)
	if len(lastSeenTxid) > 0 {
		endpoint += fmt.Sprintf("/%s", lastSeenTxid[0])
	}
	var txs []Transaction
	err := c.sendRequest(endpoint, &txs)
	return txs, err
}

// GetScripthashNewestTxsStartFromTxid  Get confirmed transaction history for the specified scripthash, sorted with newest first.
// Returns 25 transactions per page. More can be requested by specifying the last txid seen by the previous query.
func (c *EsploraClient) GetScripthashNewestTxsStartFromTxid(hash string, lastSeenTxid ...string) ([]Transaction, error) {
	endpoint := fmt.Sprintf("/scripthash/%s/txs/chain", hash)
	if len(lastSeenTxid) > 0 {
		endpoint += fmt.Sprintf("/%s", lastSeenTxid[0])
	}
	var txs []Transaction
	err := c.sendRequest(endpoint, &txs)
	return txs, err
}

// GetAllAddressTxs fetches all confirmed transaction history for the specified address
// WARM: It may take a long time if there are many tx at one address.
func (c *EsploraClient) GetAllAddressTxs(address string) ([]Transaction, error) {
	var allTxs []Transaction
	var lastTxid string
	for {
		txs, err := c.GetAddressNewestTxsStartFromTxid(address, lastTxid)
		if err != nil {
			return nil, err
		}
		if len(txs) == 0 {
			break
		}
		allTxs = append(allTxs, txs...)
		lastTxid = txs[len(txs)-1].Txid
	}
	return allTxs, nil
}

// GetAllScripthashTxs fetches all confirmed transaction history for the specified scripthash
// WARM: It may take a long time if there are many tx at one address.
func (c *EsploraClient) GetAllScripthashTxs(hash string) ([]Transaction, error) {
	var allTxs []Transaction
	var lastTxid string
	for {
		txs, err := c.GetScripthashNewestTxsStartFromTxid(hash, lastTxid)
		if err != nil {
			return nil, err
		}
		if len(txs) == 0 {
			break
		}
		allTxs = append(allTxs, txs...)
		lastTxid = txs[len(txs)-1].Txid
	}
	return allTxs, nil
}

// GetAddressTxsMempool fetches unconfirmed transaction history for the specified address
func (c *EsploraClient) GetAddressTxsMempool(address string) ([]Transaction, error) {
	var txs []Transaction
	err := c.sendRequest(fmt.Sprintf("/address/%s/txs/mempool", address), &txs)
	return txs, err
}

// GetScripthashTxsMempool fetches unconfirmed transaction history for the specified scripthash
func (c *EsploraClient) GetScripthashTxsMempool(hash string) ([]Transaction, error) {
	var txs []Transaction
	err := c.sendRequest(fmt.Sprintf("/scripthash/%s/txs/mempool", hash), &txs)
	return txs, err
}

// GetAddressUtxos fetches the list of unspent transaction outputs associated with the address
func (c *EsploraClient) GetAddressUtxos(address string) ([]Utxo, error) {
	var utxos []Utxo
	err := c.sendRequest(fmt.Sprintf("/address/%s/utxo", address), &utxos)
	return utxos, err
}

// GetScripthashUtxos fetches the list of unspent transaction outputs associated with the scripthash
func (c *EsploraClient) GetScripthashUtxos(hash string) ([]Utxo, error) {
	var utxos []Utxo
	err := c.sendRequest(fmt.Sprintf("/scripthash/%s/utxo", hash), &utxos)
	return utxos, err
}

// SearchAddressesByPrefix searches for addresses beginning with the specified prefix
// Returns a array with up to 10 results.
func (c *EsploraClient) SearchAddressesByPrefix(prefix string) ([]string, error) {
	var addresses []string
	err := c.sendRequest(fmt.Sprintf("/address-prefix/%s", prefix), &addresses)
	return addresses, err
}
