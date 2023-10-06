# TRON SDK

The **TRON SDK** is a Go client library designed to simplify interaction with the TRON blockchain's API, specifically the **TronGrid** service. TronGrid is a public API service provided by the TRON Foundation that offers various endpoints for querying TRON blockchain data.

## Features

- Retrieve information about blocks, transactions, and contract events.
- Check account balances for native tokens and TRC20 tokens.
- Obtain details about transactions and their results.
- Interact with smart contract functions and retrieve contract owner information.

## Installation

To install the **tron-api** SDK, you can use the `go get` command:

```bash
go get -u github.com/distributed-lab/tron-api
```

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/distributed-lab/tron-sdk/tron_api"
)

func main() {
	apiKey := "your-api-key"
	httpApiURL := "https://api.trongrid.io"
	grpcURL := "grpc.trongrid.io:50051"

	client := tron_api.NewTronClient(httpApiURL, grpcURL, apiKey)

	blockNumber, err := client.GetNowBlock()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Latest Block Number:", blockNumber)

	// Perform other operations using the provided methods
	// ...
}
```

## TronGrid Usage

This SDK interacts with the TRON blockchain through the **TronGrid** service, which offers a range of methods to retrieve blockchain data. Some of the key methods utilized by this SDK include:

- `GetNowBlock()`: Retrieve the latest block number.
- `GetTransactionInfoByBlockNum(block int64)`: Get transaction information by block number.
- `GetTransactionResultById(id string)`: Fetch the result of a transaction by its ID.
- `GetTxInfo(id string)`: Obtain detailed transaction information by ID.
- `BalanceOf(address address.Address, tokenAddress string)`: Check account balances for native tokens and TRC20 tokens.
- `GetContractOwner(c address.Address)`: Retrieve the owner of a smart contract.

For more details on these methods and other available endpoints, refer to the official [TronGrid API documentation](https://developers.tron.network/reference/full-node-api-overview).