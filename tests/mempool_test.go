package tests

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestGetMempoolStats(t *testing.T) {
	stats, err := realClient.GetMempoolInfo()
	assert.NoError(t, err)
	assert.NotNil(t, stats)
	assert.Greater(t, stats.Count, 0)
	assert.Greater(t, stats.Vsize, 0)
	assert.Greater(t, stats.TotalFee, 0)
	assert.NotEmpty(t, stats.FeeHistogram)
	spew.Dump(stats)
}
	

func TestGetMempoolTxids(t *testing.T) {
	txids, err := realClient.GetMempoolTxids()
	assert.NoError(t, err)
	assert.NotEmpty(t, txids)
	spew.Dump(txids)
}
	
func TestGetRecentMempoolTxs(t *testing.T) {
	txs, err := realClient.GetRecentMempoolTxs()
	assert.NoError(t, err)
	assert.NotEmpty(t, txs)
		spew.Dump(txs)
}
