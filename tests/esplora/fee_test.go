package esplora_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestGetFee(t *testing.T) {
	res, err := realClient.GetFeeEstimates()
	assert.NoError(t, err)
	spew.Dump(res)
}
