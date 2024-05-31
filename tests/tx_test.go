package tests

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactionHex(t *testing.T) {
	hex, err := realClient.GetTransactionHex("7a8873d7115b374f214d3860056c5a517793535f02387edd25b87c89f714a3e6")
	assert.NoError(t, err)
	assert.NotEmpty(t, hex)
	spew.Dump(hex)
}

func TestGetTransactionRaw(t *testing.T) {
	raw, err := realClient.GetTransactionRaw("7a8873d7115b374f214d3860056c5a517793535f02387edd25b87c89f714a3e6")
	assert.NoError(t, err)
	assert.NotEmpty(t, raw)
	spew.Dump(raw)
}

func TestGetTransactionMerkleBlockProof(t *testing.T) {
	proof, err := realClient.GetTransactionMerkleBlockProof("7a8873d7115b374f214d3860056c5a517793535f02387edd25b87c89f714a3e6")
	assert.NoError(t, err)
	assert.NotNil(t, proof)
	spew.Dump(proof)
}

func TestGetTransactionMerkleProof(t *testing.T) {
	proof, err := realClient.GetTransactionMerkleProof("7a8873d7115b374f214d3860056c5a517793535f02387edd25b87c89f714a3e6")
	assert.NoError(t, err)
	assert.NotNil(t, proof)
	spew.Dump(proof)
}

func TestGetTransactionOutputSpendingStatus(t *testing.T) {
	status, err := realClient.GetTransactionOutputSpendingStatus("7a8873d7115b374f214d3860056c5a517793535f02387edd25b87c89f714a3e6", 0)
	assert.NoError(t, err)
	assert.NotNil(t, status)
	spew.Dump(status)
}

func TestGetTransactionOutputsSpendingStatus(t *testing.T) {
	status, err := realClient.GetTransactionOutputsSpendingStatus("7a8873d7115b374f214d3860056c5a517793535f02387edd25b87c89f714a3e6")
	assert.NoError(t, err)
	assert.NotEmpty(t, status)
	spew.Dump(status)
}

func TestBroadcastTransaction(t *testing.T) {
	rawTx := "02000000000101fb4cfc16a5b31a6ee7cf3945b8337aeb5eea21135b22c56e44b9d6a65c430b210100000000fdffffff02102700000000000022512008c45a8cb08537d6719e0a4729b789414d92b8996cb730e6d94378fededebf51041eed0500000000225120bf064d51aac363f4905de2f561a546346c8b016afaaacc5119a075cca862439d0140c6b66be5669a6b00911b963c3e4571b9575d597cbbbe36e683364e10a6d59e5fb8c6af16a1bd157235eb8e5943d57b1bc712685977506aec3cd24dba073a8cd39e032b00"
	txid, err := realClient.BroadcastTransaction(rawTx)
	assert.NoError(t, err)
	assert.NotEmpty(t, txid)
	spew.Dump(txid)
}
