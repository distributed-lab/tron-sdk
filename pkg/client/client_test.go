package client_test

import (
	"testing"
	"time"

	"github.com/distributed-lab/tron-sdk/pkg/client"
	"github.com/stretchr/testify/assert"
)

const (
	apiKey  = ""
	httpApi = ""
	grpcUrl = ""
)

func TestGetNowBlock(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		tronClient := client.NewTronClient(httpApi, grpcUrl, apiKey)

		blockNumber, err := tronClient.GetNowBlock()
		assert.NoError(t, err)
		assert.NotEqual(t, int64(0), blockNumber)
	})
}

func TestGetTransactionInfoByBlockNum(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		tronClient := client.NewTronClient(httpApi, grpcUrl, apiKey)

		blockNumber := int64(1)
		txs, blockTime, err := tronClient.GetTransactionInfoByBlockNum(blockNumber)
		assert.NoError(t, err)
		assert.NotNil(t, txs)
		assert.Equal(t, time.Time{}, blockTime)
	})
}

func TestGetContractOwner(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		tronClient := client.NewTronClient(httpApi, grpcUrl, apiKey)
		contractAddress := "asdadasdasdasdas"
		owner, err := tronClient.GetContractOwner([]byte(contractAddress))
		assert.Error(t, err, "failed to get contract owner: invalid address length: 46")
		assert.Nil(t, owner)
	})

}

func TestGetBlock(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		tronClient := client.NewTronClient(httpApi, grpcUrl, apiKey)
		block, err := tronClient.GetBlock("1", false)

		assert.NoError(t, err)
		assert.Equal(t, block.BlockID, "0000000000000001049f911bc1069bfd2c2225bc3cd210abd02fb219751813f0")
		assert.Equal(t, block.BlockHeader.RawData.Number, int64(1))
		assert.Equal(t, block.BlockHeader.RawData.ParentHash, "0000000000000000de1aa88295e1fcf982742f773e0419c5a9c134c994a9059e")
	})

}

func TestGetTransactionInfoByContractAddress(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		tronClient := client.NewTronClient(httpApi, grpcUrl, apiKey)
		txs, err := tronClient.GetTxInfoByContractAddress("TG3XXyExBkPp9nzdajDZsozEu4BkaSJozs")

		assert.NoError(t, err)
		assert.NotNil(t, txs.Data[0].BlockNumber)
	})

}

func TestGetEventsByTxID(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		tronClient := client.NewTronClient(httpApi, grpcUrl, apiKey)
		events, err := tronClient.GetEventsByTxID(
			"1171e17ec5e1847f7e3938b574cfc06a397ab5064e839c4a55bacd12dca987e6",
		)

		assert.NoError(t, err)
		assert.True(t, events.Success)
	})
}

func TestGetTxInfoByAccountAddress(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		tronClient := client.NewTronClient(httpApi, grpcUrl, apiKey)
		_, err := tronClient.GetTxInfoByAccountAddress(
			"TSvu34UfFf3f73NPrAmUH64VdQNjCsN7YL",
		)

		assert.NoError(t, err)
	})
}
