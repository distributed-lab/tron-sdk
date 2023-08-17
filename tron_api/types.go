package tron_api

const (
	TronSucceededTx = "SUCCESS"
	TronRevertedTx  = "REVERT"
	BalanceOfPrefix = "70a08231000000000000000000000000"
)

const NativeToken = "T9yD14Nj9j7xAB4dbGeiX9h8unkKHxuWwb"

type Block struct {
	BlockID     string `json:"blockID"`
	BlockHeader struct {
		RawData struct {
			Number         int64  `json:"number"`
			TxTrieRoot     string `json:"txTrieRoot"`
			WitnessAddress string `json:"witness_address"`
			ParentHash     string `json:"parentHash"`
			Version        int64  `json:"version"`
			Timestamp      int64  `json:"timestamp"`
		} `json:"raw_data"`
		WitnessSignature string `json:"witness_signature"`
	} `json:"block_header"`
	Transactions []struct {
		Ret []struct {
			ContractRet string `json:"contractRet"`
		} `json:"ret"`
		Signature []string `json:"signature"`
		TxID      string   `json:"txID"`
		RawData   struct {
			Contract []struct {
				Parameter struct {
					Value struct {
						Amount       int64  `json:"amount"`
						OwnerAddress string `json:"owner_address"`
						ToAddress    string `json:"to_address"`
					} `json:"value"`
					TypeUrl string `json:"type_url"`
				} `json:"parameter"`
				Type string `json:"type"`
			} `json:"contract"`
			RefBlockBytes string `json:"ref_block_bytes"`
			RefBlockHash  string `json:"ref_block_hash"`
			Expiration    int64  `json:"expiration"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"raw_data"`
		RawDataHex string `json:"raw_data_hex"`
	} `json:"transactions"`
}

type TxInfo struct {
	Id              string   `json:"id"`
	Fee             int64    `json:"fee"`
	BlockNumber     int64    `json:"blockNumber"`
	BlockTimestamp  int64    `json:"blockTimeStamp"`
	ContractResult  []string `json:"contractResult"`
	ContractAddress string   `json:"contract_address"`
	Receipt         struct {
		EnergyFee        int64  `json:"energy_fee"`
		EnergyUsageTotal int64  `json:"energy_usage_total"`
		NetFee           int64  `json:"net_fee"`
		Result           string `json:"result"`
	} `json:"receipt"`
	Log []struct {
		Address string   `json:"address"`
		Topics  []string `json:"topics"`
		Data    string   `json:"data"`
	} `json:"log"`
	InternalTransactions []struct {
		Hash              string `json:"hash"`
		CallerAddress     string `json:"caller_address"`
		TransferToAddress string `json:"transferTo_address"`
		CallValueInfo     []struct {
			CallValue int64 `json:"callValue"`
		} `json:"callValueInfo"`
		Note string `json:"note"`
	} `json:"internal_transactions"`
	PackingFee int64 `json:"packingFee"`
}

type Tx struct {
	Ret []struct {
		ContractRet string `json:"contractRet"`
	} `json:"ret"`
	RawData struct {
		Contract []struct {
			Type      string `json:"type"`
			Parameter struct {
				Value struct {
					Amount       int64  `json:"amount"`
					OwnerAddress string `json:"owner_address"`
					ToAddress    string `json:"to_address"`
				} `json:"value"`
			} `json:"parameter"`
		} `json:"contract"`
	} `json:"raw_data"`
	ContractAddress string `json:"contract_address"`
}

type Balance struct {
	ConstantResult []string `json:"constant_result"`
}
