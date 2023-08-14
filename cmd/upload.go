package cmd

import (
	"github.com/Conflux-Chain/neurahive-client/common/blockchain"
	"github.com/Conflux-Chain/neurahive-client/contract"
	"github.com/Conflux-Chain/neurahive-client/file"
	"github.com/Conflux-Chain/neurahive-client/node"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	uploadArgs struct {
		file string
		tags string

		url      string
		contract string
		key      string

		node string

		force bool
	}

	uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "Upload file to Neurahive network",
		Run:   upload,
	}
)

func init() {
	uploadCmd.Flags().StringVar(&uploadArgs.file, "file", "", "File name to upload")
	uploadCmd.MarkFlagRequired("file")
	uploadCmd.Flags().StringVar(&uploadArgs.tags, "tags", "0x", "Tags of the file")

	uploadCmd.Flags().StringVar(&uploadArgs.url, "url", "", "Fullnode URL to interact with Neurahive smart contract")
	uploadCmd.MarkFlagRequired("url")
	uploadCmd.Flags().StringVar(&uploadArgs.contract, "contract", "", "Neurahive smart contract to interact with")
	uploadCmd.MarkFlagRequired("contract")
	uploadCmd.Flags().StringVar(&uploadArgs.key, "key", "", "Private key to interact with smart contract")
	uploadCmd.MarkFlagRequired("key")

	uploadCmd.Flags().StringVar(&uploadArgs.node, "node", "", "Neurahive storage node URL")
	uploadCmd.MarkFlagRequired("node")

	uploadCmd.Flags().BoolVar(&uploadArgs.force, "force", false, "Force to upload file even already exists")

	rootCmd.AddCommand(uploadCmd)
}

func upload(*cobra.Command, []string) {
	client := blockchain.MustNewWeb3(uploadArgs.url, uploadArgs.key)
	defer client.Close()
	contractAddr := common.HexToAddress(uploadArgs.contract)
	flow, err := contract.NewFlowContract(contractAddr, client)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to create flow contract")
	}

	node := node.MustNewClient(uploadArgs.node)
	defer node.Close()

	uploader := file.NewUploader(flow, node)
	opt := file.UploadOption{
		Tags:  hexutil.MustDecode(uploadArgs.tags),
		Force: uploadArgs.force,
	}
	if err := uploader.Upload(uploadArgs.file, opt); err != nil {
		logrus.WithError(err).Fatal("Failed to upload file")
	}
}
