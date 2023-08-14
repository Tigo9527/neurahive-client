package main

import (
	"fmt"

	"github.com/Conflux-Chain/neurahive-client/common/blockchain"
	"github.com/Conflux-Chain/neurahive-client/contract"
	"github.com/Conflux-Chain/neurahive-client/kv"
	"github.com/Conflux-Chain/neurahive-client/node"
	"github.com/ethereum/go-ethereum/common"
)

const NrhvClientAddr = "http://127.0.0.1:5678"
const BlockchainClientAddr = ""
const PrivKey = ""
const FlowContractAddr = ""

func main() {
	nrhvClient, err := node.NewClient(NrhvClientAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	blockchainClient := blockchain.MustNewWeb3(BlockchainClientAddr, PrivKey)
	defer blockchainClient.Close()
	blockchain.CustomGasLimit = 1000000
	nrhv, err := contract.NewFlowContract(common.HexToAddress(FlowContractAddr), blockchainClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	kvClient := kv.NewClient(nrhvClient, nrhv)
	batcher := kvClient.Batcher()
	batcher.Set(common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000f2bd"),
		[]byte("TESTKEY0"),
		[]byte{69, 70, 71, 72, 73},
	)
	batcher.Set(common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000f2bd"),
		[]byte("TESTKEY1"),
		[]byte{74, 75, 76, 77, 78},
	)
	err = batcher.Exec()
	if err != nil {
		fmt.Println(err)
		return
	}
}
