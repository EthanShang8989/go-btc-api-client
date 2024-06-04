package mempoolspace_test

import (
	"testing"
	"time"

	"github.com/EthanShang8989/go-btc-api-client/btcclient"
	"github.com/EthanShang8989/go-btc-api-client/btcclient/mempoolspace"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

var realClient = mempoolspace.NewClient(btcclient.MenpoolTestnetURL)

func TestGetBtcPrice(t *testing.T) {
	var realClient = mempoolspace.NewClient(btcclient.MempoolURL)
	res, err := realClient.GetBtcPrice()
	assert.NoError(t, err)
	assert.True(t, res.Time > 1000)
	spew.Dump(res)
}
func TestGetDifficultyAdjustment(t *testing.T) {
	res, err := realClient.GetDifficultyAdjustment()
	assert.NoError(t, err)
	assert.True(t, res.AdjustedTimeAvg > 10)
	spew.Dump(res)
}
func TestGetHistoricalPrice(t *testing.T) {
	var realClient = mempoolspace.NewClient(btcclient.MempoolURL)
	res, err := realClient.GetHistoricalPrice("EUR", time.Now().Unix())
	assert.NoError(t, err)
	assert.True(t, res.Prices[0].Time > 1000)
	// spew.Dump(res)
}
