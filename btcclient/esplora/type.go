package esplora

import (
	"encoding/json"
	"errors"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

type BlockHeader struct {
	ID                string         `json:"id"`
	Height            uint64         `json:"height"`
	Version           int            `json:"version"`
	Timestamp         int64          `json:"timestamp"`
	Bits              int            `json:"bits"`
	Nonce             int64          `json:"nonce"`
	Difficulty        float64        `json:"difficulty"`
	MerkleRoot        chainhash.Hash `json:"merkle_root"`
	TxCount           uint64         `json:"tx_count"`
	Size              int            `json:"size"`
	Weight            int            `json:"weight"`
	PreviousBlockHash chainhash.Hash `json:"previousblockhash"`
	Mediantime        int64          `json:"mediantime"`
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

	//Elements only
	IsPegin  bool      `json:"is_pegin"`
	Issuance *Issuance `json:"issuance"`
}

// Issuance represents issuance information for an Elements transaction.
type Issuance struct {
	AssetID               string  `json:"asset_id"`
	IsReissuance          bool    `json:"is_reissuance"`
	AssetBlindingNonce    string  `json:"asset_blinding_nonce"`
	AssetEntropy          string  `json:"asset_entropy"`
	ContractHash          string  `json:"contract_hash"`
	AssetAmount           float64 `json:"assetamount,omitempty"`
	AssetAmountCommitment string  `json:"assetamountcommitment,omitempty"`
	TokenAmount           float64 `json:"tokenamount,omitempty"`
	TokenAmountCommitment string  `json:"tokenamountcommitment,omitempty"`
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

// UserIssuedAsset represents an asset in the Liquid network.
type UserIssuedAsset struct {
	AssetID         string               `json:"asset_id"`
	IssuanceTxin    IssuanceTxin         `json:"issuance_txin,omitempty"`    //the issuance transaction input
	IssuancePrevout IssuancePrevout      `json:"issuance_prevout,omitempty"` //the previous output spent for the issuance
	ReissuanceToken string               `json:"reissuance_token"`           //the asset id of the reissuance token
	ContractHash    string               `json:"contract_hash"`              //the contract hash committed as the issuance entropy
	Status          Status               `json:"status"`                     //the confirmation status of the initial asset issuance transaction
	ChainStats      UserIssuedAssetStats `json:"chain_stats"`
	MempoolStats    UserIssuedAssetStats `json:"mempool_stats"`
	Contract        Contract             `json:"contract"`
	Entity          Entity               `json:"entity"`    //the entity linked to this asset.
	Precision       int                  `json:"precision"` //the number of decimal places for units of this asset (defaults to 0)
	Name            string               `json:"name"`      //a description for the asset (up to 255 characters)
	Ticker          string               `json:"ticker"`    //a 3-5 characters ticker associated with the asset (optional)
}

type UserIssuedAssetStats struct {
	TxCount                int   `json:"tx_count"`                 //the number of transactions associated with this asset (does not include confidential transactions)
	IssuanceCount          int   `json:"issuance_count"`           //the number of (re)issuance transactions
	IssuedAmount           int   `json:"issued_amount"`            //the number of (re)issuance transactions
	BurnedAmount           int64 `json:"burned_amount"`            //the total amount provably burned
	HasBlindedIssuances    bool  `json:"has_blinded_issuances"`    //whether at least one of the (re)issuances were blind
	ReissuanceTokens       int   `json:"reissuance_tokens"`        //the number of reissuance tokens
	BurnedReissuanceTokens int64 `json:"burned_reissuance_tokens"` //the number of reissuance tokens burned
}

type IssuanceTxin struct {
	Txid string `json:"txid"`
	Vin  int    `json:"vin"`
}

type IssuancePrevout struct {
	Txid string `json:"txid"`
	Vout int    `json:"vout"`
}

type Contract struct {
	Entity       Entity `json:"entity"`
	IssuerPubkey string `json:"issuer_pubkey"`
	Name         string `json:"name"`
	Precision    int    `json:"precision"`
	Ticker       string `json:"ticker"`
	Version      int    `json:"version"`
}

type Entity struct {
	Domain string `json:"domain"`
}

type NativeAsset struct {
	AssetID      string           `json:"asset_id"`
	ChainStats   NativeAssetStats `json:"chain_stats"`
	MempoolStats NativeAssetStats `json:"mempool_stats"`
}
type NativeAssetStats struct {
	TxCount      int64 `json:"tx_count"`
	PegInCount   int64 `json:"peg_in_count"`
	PegInAmount  int64 `json:"peg_in_amount"`
	PegOutAmount int64 `json:"peg_out_amount"`
	PegOutCount  int64 `json:"peg_out_count"`
	BurnCount    int64 `json:"burn_count"`
	BurnedAmount int64 `json:"burned_amount"`
}
