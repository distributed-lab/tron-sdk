package tx

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/protobuf/proto"
)

func SignTx(tx *api.TransactionExtention, privateKey *ecdsa.PrivateKey) (*api.TransactionExtention, error) {
	rawData, err := proto.Marshal(tx.Transaction.GetRawData())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get rawData")
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create signature")
	}

	tx.Transaction.Signature = append(tx.Transaction.Signature, signature)

	return tx, nil
}

func SignAndBroadcastTx(
	tx *api.TransactionExtention, privateKey *ecdsa.PrivateKey, cli *client.GrpcClient,
) (*core.TransactionInfo, error) {
	signedTx, err := SignTx(tx, privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign tx")
	}

	result, err := cli.Broadcast(signedTx.Transaction)
	if err != nil {
		return nil, errors.Wrap(err, "failed to broadcast tx")
	}

	if !result.GetResult() {
		return nil, errors.New("error tx result")
	}

	txInfo, err := WaitMined(cli, hex.EncodeToString(signedTx.Txid))
	if err != nil {
		return nil, errors.Wrap(err, "failed to wait mined", logan.F{
			"tx_id": hex.EncodeToString(signedTx.Txid),
		})
	}

	return txInfo, nil
}
