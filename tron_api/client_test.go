package tron_api_test

import (
	"testing"
	"time"

	"github.com/distributed-lab/tron-sdk/tron_api"
	"github.com/stretchr/testify/assert"
)

const (
	apiKey  = ""
	httpApi = ""
	grpcUrl = ""
)

func TestGetNowBlock(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		tronClient := tron_api.NewTronClient(httpApi, grpcUrl, apiKey)

		blockNumber, err := tronClient.GetNowBlock()
		assert.NoError(t, err)
		assert.NotEqual(t, int64(0), blockNumber)
	})
}

func TestGetTransactionInfoByBlockNum(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		tronClient := tron_api.NewTronClient(httpApi, grpcUrl, apiKey)

		blockNumber := int64(1)
		txs, blockTime, err := tronClient.GetTransactionInfoByBlockNum(blockNumber)
		assert.NoError(t, err)
		assert.NotNil(t, txs)
		assert.Equal(t, time.Time{}, blockTime)
	})
}

func TestGetContractOwner(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		tronClient := tron_api.NewTronClient(httpApi, grpcUrl, apiKey)
		contractAddress := "asdadasdasdasdas"
		owner, err := tronClient.GetContractOwner([]byte(contractAddress))
		assert.Error(t, err, "failed to get contract owner: invalid address length: 46")
		assert.Nil(t, owner)
	})

}
