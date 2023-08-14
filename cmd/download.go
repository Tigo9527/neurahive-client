package cmd

import (
	"github.com/Conflux-Chain/neurahive-client/file"
	"github.com/Conflux-Chain/neurahive-client/node"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	downloadArgs struct {
		file  string
		nodes []string
		root  string
		proof bool
	}

	downloadCmd = &cobra.Command{
		Use:   "download",
		Short: "Download file from Neurahive network",
		Run:   download,
	}
)

func init() {
	downloadCmd.Flags().StringVar(&downloadArgs.file, "file", "", "File name to download")
	downloadCmd.MarkFlagRequired("file")
	downloadCmd.Flags().StringSliceVar(&downloadArgs.nodes, "node", []string{}, "Neurahive storage node URL. Multiple nodes could be specified and separated by comma, e.g. url1,url2,url3")
	downloadCmd.MarkFlagRequired("node")
	downloadCmd.Flags().StringVar(&downloadArgs.root, "root", "", "Merkle root to download file")
	downloadCmd.MarkFlagRequired("root")
	downloadCmd.Flags().BoolVar(&downloadArgs.proof, "proof", false, "Whether to download with merkle proof for validation")

	rootCmd.AddCommand(downloadCmd)
}

func download(*cobra.Command, []string) {
	nodes := node.MustNewClients(downloadArgs.nodes)

	downloader := file.NewDownloader(nodes...)

	if err := downloader.Download(downloadArgs.root, downloadArgs.file, downloadArgs.proof); err != nil {
		logrus.WithError(err).Fatal("Failed to download file")
	}
}
