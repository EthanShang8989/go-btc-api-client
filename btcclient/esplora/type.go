package esplora

import (
	"encoding/json"
	"errors"
)

type BlockHeader struct {
	ID                string  `json:"id"`
	Height            uint64  `json:"height"`
	Version           int     `json:"version"`
	Timestamp         int64   `json:"timestamp"`
	Bits              int     `json:"bits"`
	Nonce             int64   `json:"nonce"`
	Difficulty        float64 `json:"difficulty"`
	MerkleRoot        string  `json:"merkle_root"`
	TxCount           uint64  `json:"tx_count"`
	Size              int     `json:"size"`
	Weight            int     `json:"weight"`
	PreviousBlockHash string  `json:"previousblockhash"`
	Mediantime        int64   `json:"mediantime"`
}
type BlockStatus struct {
	InBestChain bool   `json:"in_best_chain"`
	Height      uint64 `json:"height"`
	NextBest    string `json:"next_best"`
}

type Transaction struct {
	Txid     string `json:"txid"`
	Version  uint64 `json:"version"`
	Locktime uint64 `json:"locktime"`
	Size     uint64 `json:"size"`
	Weight   uint64 `json:"weight"`
	Fee      uint64 `json:"fee"`
	Vin      []Vin  `json:"vin"`
	Vout     []Vout `json:"vout"`
	Status   Status `json:"status"`
}

type Vin struct {
	TxID                  string   `json:"txid"`
	Vout                  int      `json:"vout"`
	IsCoinbase            bool     `json:"is_coinbase"`
	ScriptSig             string   `json:"scriptsig,omitempty"`
	ScriptSigAsm          string   `json:"scriptsig_asm,omitempty"`
	InnerRedeemScriptAsm  string   `json:"inner_redeemscript_asm,omitempty"`
	InnerWitnessScriptAsm string   `json:"inner_witnessscript_asm,omitempty"`
	Sequence              int      `json:"sequence"`
	Witness               []string `json:"witness"`
	Prevout               *Vout    `json:"prevout"`
}

type Vout struct {
	Scriptpubkey        string `json:"scriptpubkey,omitempty"`
	ScriptpubkeyAsm     string `json:"scriptpubkey_asm,omitempty"`
	ScriptpubkeyType    string `json:"scriptpubkey_type,omitempty"`
	ScriptpubkeyAddress string `json:"scriptpubkey_address,omitempty"`
	Value               uint64 `json:"value"`
}

// Status 表示tx或者utxo的状态
type Status struct {
	Confirmed   bool   `json:"confirmed"`
	BlockHeight int    `json:"block_height"`
	BlockHash   string `json:"block_hash"`
	BlockTime   int64  `json:"block_time"`
}

type AddressInfo struct {
	Address       string `json:"address"`
	MempoolStatus `json:"mempool_stats"`
	ChainStats    `json:"chain_stats"`
}
type ScripthashInfo struct {
	MempoolStatus `json:"mempool_stats"`
	ChainStats    `json:"chain_stats"`
	Scripthash    string `json:"scripthash"`
}

type ChainStats struct {
	FundedTxoCount int64 `json:"funded_txo_count"`
	FundedTxoSum   int64 `json:"funded_txo_sum"`
	SpentTxoCount  int64 `json:"spent_txo_count"`
	SpentTxoSum    int64 `json:"spent_txo_sum"`
	TxCount        int64 `json:"tx_count"`
}

type MempoolStatus struct {
	FundedTxoCount int64 `json:"funded_txo_count"`
	FundedTxoSum   int64 `json:"funded_txo_sum"`
	SpentTxoCount  int64 `json:"spent_txo_count"`
	SpentTxoSum    int64 `json:"spent_txo_sum"`
	TxCount        int64 `json:"tx_count"`
}

type Utxo struct {
	TxID   string `json:"txid"`
	Vout   int    `json:"vout"`
	Value  uint64 `json:"value"`
	Status `json:"status"`
}

type MempoolInfo struct {
	Count        int            `json:"count"`
	Vsize        int            `json:"vsize"`
	TotalFee     int            `json:"total_fee"`
	FeeHistogram []FeeHistogram `json:"fee_histogram"`
}

// FeeHistogram represents a fee rate and corresponding virtual size
type FeeHistogram struct {
	Feerate float64
	Vsize   float64
}

// UnmarshalJSON custom unmarshals a FeeHistogram from a JSON array
func (fh *FeeHistogram) UnmarshalJSON(data []byte) error {
	var tuple [2]float64
	if err := json.Unmarshal(data, &tuple); err != nil {
		return err
	}
	if len(tuple) != 2 {
		return errors.New("invalid tuple length")
	}
	fh.Feerate = tuple[0]
	fh.Vsize = tuple[1]
	return nil
}

type MerkleProof struct {
	BlockHeight int      `json:"block_height"`
	Merkle      []string `json:"merkle"`
	Pos         int      `json:"pos"`
}
