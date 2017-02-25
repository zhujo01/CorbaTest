package libcluster

import (
	"io"
	"github.com/docker/docker/daemon/cluster"
	"github.com/zhujo01/CorbaTest/libcluster/persist"
)

type API interface {
	io.Closer
	NewCluster(driverName string) (*cluster.Cluster, error)
	Create(c *cluster.Cluster) error
	persist.Store
	GetClustersDir() string
}

type Client struct {
	*persist.Filestore
}

func NewClient(storPath, certsDir string) *Client {
	return &Client{}
}

func (c *Client) NewCluster(driverName string) (*cluster.Cluster, error) {
	return nil
}

func (c *Client) Close() error {
	return nil
}

func (c *Client) Create(c *cluster.Cluster) error {
	return nil
}

func (c *Client) GetClustersDir() string {
	return nil
}

