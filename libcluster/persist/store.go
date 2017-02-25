package persist

import "github.com/docker/docker/daemon/cluster"

type Store interface {
	// Exists returns whether a cluster exists or not
	Exists(name string) (bool, error)

	// List returns a list of all clusters in the store
	List() ([]string, error)

	// Load loads a cluster by name
	Load(name string) (*cluster.Cluster, error)

	// Remove removes a machine from the store
	Remove(name string) error

	// Save persists a cluster in the store
	Save(host *cluster.Cluster) error
}
