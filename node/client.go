package node

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/sirupsen/logrus"
)

type Client struct {
	url string
	*providers.MiddlewarableProvider

	nrhv  *NeurahiveClient
	admin *AdminClient
	kv    *KvClient
}

func MustNewClient(url string, option ...providers.Option) *Client {
	client, err := NewClient(url, option...)
	if err != nil {
		logrus.WithError(err).WithField("url", url).Fatal("Failed to connect to storage node")
	}

	return client
}

func NewClient(url string, option ...providers.Option) (*Client, error) {
	var opt providers.Option
	if len(option) > 0 {
		opt = option[0]
	}

	provider, err := providers.NewProviderWithOption(url, opt)
	if err != nil {
		return nil, err
	}

	return &Client{
		url:                   url,
		MiddlewarableProvider: provider,

		nrhv:  &NeurahiveClient{provider},
		admin: &AdminClient{provider},
		kv:    &KvClient{provider},
	}, nil
}

func MustNewClients(urls []string, option ...providers.Option) []*Client {
	var clients []*Client

	for _, url := range urls {
		client := MustNewClient(url, option...)
		clients = append(clients, client)
	}

	return clients
}

func (c *Client) URL() string {
	return c.url
}

func (c *Client) Neurahive() *NeurahiveClient {
	return c.nrhv
}

func (c *Client) Admin() *AdminClient {
	return c.admin
}

func (c *Client) KV() *KvClient {
	return c.kv
}

// Neurahive RPCs
type NeurahiveClient struct {
	provider *providers.MiddlewarableProvider
}

func (c *NeurahiveClient) GetStatus() (status Status, err error) {
	err = c.provider.CallContext(context.Background(), &status, "nrhv_getStatus")
	return
}

func (c *NeurahiveClient) GetFileInfo(root common.Hash) (file *FileInfo, err error) {
	err = c.provider.CallContext(context.Background(), &file, "nrhv_getFileInfo", root)
	return
}

func (c *NeurahiveClient) GetFileInfoByTxSeq(txSeq uint64) (file *FileInfo, err error) {
	err = c.provider.CallContext(context.Background(), &file, "nrhv_getFileInfoByTxSeq", txSeq)
	return
}

func (c *NeurahiveClient) UploadSegment(segment SegmentWithProof) (ret int, err error) {
	err = c.provider.CallContext(context.Background(), &ret, "nrhv_uploadSegment", segment)
	return
}

func (c *NeurahiveClient) DownloadSegment(root common.Hash, startIndex, endIndex uint64) (data []byte, err error) {
	err = c.provider.CallContext(context.Background(), &data, "nrhv_downloadSegment", root, startIndex, endIndex)
	return
}

func (c *NeurahiveClient) DownloadSegmentWithProof(root common.Hash, index uint64) (segment *SegmentWithProof, err error) {
	err = c.provider.CallContext(context.Background(), &segment, "nrhv_downloadSegmentWithProof", root, index)
	return
}

// Admin RPCs
type AdminClient struct {
	provider *providers.MiddlewarableProvider
}

func (c *AdminClient) Shutdown() (ret int, err error) {
	err = c.provider.CallContext(context.Background(), &ret, "admin_shutdown")
	return
}

func (c *AdminClient) StartSyncFile(txSeq uint64) (ret int, err error) {
	err = c.provider.CallContext(context.Background(), &ret, "admin_startSyncFile", txSeq)
	return
}

func (c *AdminClient) GetSyncStatus(txSeq uint64) (status string, err error) {
	err = c.provider.CallContext(context.Background(), &status, "admin_getSyncStatus", txSeq)
	return
}
