package helpers

import (
	"crypto/ecdsa"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewTransferTrxTx(
	c *client.GrpcClient, addressFrom, addressTo string, amount int64, signatures [][]byte, privateKeys ...*ecdsa.PrivateKey,
) (*api.TransactionExtention, error) {
	account, err := c.GetAccount(addressFrom)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account")
	}

	if account.GetBalance() < amount {
		return nil, ErrInsufficientFunds
	}

	tx, err := c.Transfer(addressFrom, addressTo, amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create tx Transfer TRX")
	}

	tx, err = AddSignaturesAndBroadcast(c, tx, signatures, privateKeys...)

	if err != nil && tx.Result == nil {
		return nil, errors.Wrap(err, "failed to broadcast tx")
	}

	return tx, nil
}
