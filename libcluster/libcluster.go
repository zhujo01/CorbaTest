package libcluster

import (
	"io"
	"github.com/zhujo01/CorbaTest/libcluster/cluster"
	"github.com/zhujo01/CorbaTest/libcluster/persist"
)

type API interface {
	io.Closer
	NewCluster(cloudDriverName, clusterDriverName string) (*cluster.Cluster, error)
	Create(c *cluster.Cluster) error
	persist.Store
	GetClustersDir() string
}

type Client struct {
	// image ids goes here...
	*persist.Filestore
}

func NewClient(storPath, certsDir string) *Client {
	return &Client{}
}

func (c *Client) NewCluster(cloudDriverName, clusterDriverName string) (*cluster.Cluster, error) {
	return nil
}

func (c *Client) Close() error {
	return nil
}

func (c *Client) Create(c *cluster.Cluster) error {

	// call cloud driver to prepare (upload) images

	// call cloud driver to start storage VMs

	// call cloud driver to create docker cluster

	// configure docker swarm on docker cluster

	// create private registry in docker swarm

	// push docker images to registry

	// start app containers

	return nil
}

func (c *Client) GetClustersDir() string {
	return nil
}

