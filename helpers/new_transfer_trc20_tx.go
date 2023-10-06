package helpers

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"

	"gitlab.com/distributed_lab/logan/v3"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewTransferTrc20Tx(
	c *client.GrpcClient, addressFrom, addressTo, contract string, amount int64, signatures [][]byte, privateKeys ...*ecdsa.PrivateKey,
) (*api.TransactionExtention, error) {
	balance, err := c.TRC20ContractBalance(addressFrom, contract)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user's trc20 balance")
	}

	if balance.Cmp(big.NewInt(amount)) == -1 {
		return nil, ErrInsufficientFunds
	}

	tx, err := c.TRC20Send(addressFrom, addressTo, contract, big.NewInt(amount), 1000000000)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create tx TRC20Send")
	}

	tx, err = AddSignaturesAndBroadcast(c, tx, signatures, privateKeys...)

	if err != nil && tx.Result == nil {
		return nil, errors.Wrap(err, "failed to broadcast tx")
	}

	return tx, nil
}

func AddSignaturesAndBroadcast(
	c *client.GrpcClient, tx *api.TransactionExtention, signatures [][]byte, privateKeys ...*ecdsa.PrivateKey,
) (*api.TransactionExtention, error) {
	var err error
	if len(signatures) > 0 {
		tx.Transaction.Signature = append(tx.Transaction.Signature, signatures...)
	}

	for _, sk := range privateKeys {
		tx, err = signTx(tx, sk)
		if err != nil {
			return nil, errors.Wrap(err, "failed to sign tx", logan.F{
				"tx_id": hex.EncodeToString(tx.Txid),
			})
		}
	}

	// if err != nil and result != nil, so we need check error inside result
	result, err := c.Broadcast(tx.Transaction)
	if result != nil {
		tx.Result = result
	}

	return tx, err
}
