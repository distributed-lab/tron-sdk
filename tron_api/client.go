package tron_api

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	gotroncli "github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/martinboehm/btcutil/base58"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TronClient struct {
	apiKey     string
	rpcUrl     string
	client     *http.Client
	WrappedCli *gotroncli.GrpcClient
}

func NewTronClient(httpApi, grpcUrl, apiKey string) *TronClient {
	wrappedCli := gotroncli.NewGrpcClient(grpcUrl)
	if err := wrappedCli.Start(grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		panic(errors.Wrap(err, "failed to start grpc client"))
	}

	return &TronClient{
		apiKey:     apiKey,
		rpcUrl:     httpApi,
		client:     http.DefaultClient,
		WrappedCli: wrappedCli,
	}
}

func (cli *TronClient) GetNowBlock() (int64, error) {
	resp, err := cli.post("/wallet/getnowblock", nil)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get last block")
	}

	var block Block
	if err = cli.processResponse(resp, &block); err != nil {
		return 0, errors.Wrap(err, "failed to process response")
	}

	return block.BlockHeader.RawData.Number, nil
}

func (cli *TronClient) GetTransactionInfoByBlockNum(block int64) ([]TxInfo, time.Time, error) {
	body, err := json.Marshal(map[string]int64{
		"num": block,
	})
	if err != nil {
		return nil, time.Time{}, errors.Wrap(err, "failed to marshal params")
	}

	resp, err := cli.post("/wallet/gettransactioninfobyblocknum", body)
	if err != nil {
		return nil, time.Time{}, errors.Wrap(err, "failed to get transactions by block")
	}

	var txs []TxInfo
	if err = cli.processResponse(resp, &txs); err != nil {
		return nil, time.Time{}, errors.Wrap(err, "failed to process response")
	}

	var blockTime time.Time
	if len(txs) != 0 {
		blockTime = time.Unix(txs[0].BlockTimestamp, 0)
	}

	return txs, blockTime, nil
}

func (cli *TronClient) GetTransactionResultById(id string) (string, *core.Transaction_Result, error) {
	tx, err := cli.WrappedCli.GetTransactionByID(id)
	if err != nil && err.Error() != "transaction info not found" {
		return TronRevertedTx, nil, errors.Wrap(err, "failed to get transaction by id")
	}
	if err != nil {
		return TronRevertedTx, nil, nil
	}

	if tx.Ret[0].ContractRet.String() == TronRevertedTx {
		return TronRevertedTx, nil, nil
	}

	return TronSucceededTx, tx.Ret[0], nil
}

func (cli *TronClient) GetTxInfo(id string) (*core.TransactionInfo, error) {
	info, err := cli.WrappedCli.GetTransactionInfoByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to tx info")
	}

	return info, nil
}

func (cli *TronClient) GetTransactionById(id string) (Tx, error) {
	body, err := json.Marshal(map[string]string{
		"value": id,
	})
	if err != nil {
		return Tx{}, errors.Wrap(err, "failed to marshal params")
	}

	resp, err := cli.post("/walletsolidity/gettransactionbyid", body)
	if err != nil {
		return Tx{}, errors.Wrap(err, "failed to get transactions by block")
	}

	var tx Tx
	if err = cli.processResponse(resp, &tx); err != nil {
		return Tx{}, errors.Wrap(err, "failed to process response")
	}

	return tx, nil
}

func (cli *TronClient) BalanceOf(address address.Address, tokenAddress string) (*big.Int, error) {
	if tokenAddress == NativeToken {
		balance, err := cli.NativeBalanceOf(address.String())
		if err != nil {
			return nil, errors.Wrap(err, "failed to get balance")
		}

		return balance, nil
	}

	balance, err := cli.TRC20BalanceOf(tokenAddress, address.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get token balance")
	}

	return balance, nil
}

func (cli *TronClient) NativeBalanceOf(account string) (*big.Int, error) {
	acc, err := cli.WrappedCli.GetAccount(account)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account")
	}

	return big.NewInt(acc.Balance), nil
}

func (cli *TronClient) TRC20BalanceOf(token, account string) (*big.Int, error) {
	addr, err := address.Base58ToAddress(account)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert base58 to address")
	}
	tx, err := cli.WrappedCli.TRC20Call(
		account,
		token,
		BalanceOfPrefix+strings.TrimPrefix(addr.Hex(), "0x41"),
		true,
		0)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make a TRC20 call")
	}

	balance := new(big.Int).SetBytes(tx.GetConstantResult()[0])
	return balance, nil
}

func (cli *TronClient) GetContractOwner(c address.Address) (address.Address, error) {
	tx, err := cli.WrappedCli.TriggerConstantContract("", c.String(),
		"owner()", ``)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get contract owner")
	}

	owner, err := address.Base58ToAddress(base58.CheckEncode(
		common.HexToAddress(hex.EncodeToString(tx.ConstantResult[0])).Bytes(),
		[]byte{address.TronBytePrefix},
		base58.Sha256D))
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert base58 to address")
	}

	return owner, nil
}
