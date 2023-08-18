package tx

import (
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
)

func WaitMined(client *client.GrpcClient, id string) (*core.TransactionInfo, error) {
	for {
		logan.New().Info("wait mined tx: ", id)
		if txi, err := client.GetTransactionInfoByID(id); err == nil {
			if txi.Result != 0 {
				return txi, errors.New(string(txi.ResMessage))
			}

			return txi, nil
		}

		time.Sleep(time.Second)
	}
}
