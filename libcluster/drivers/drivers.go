package drivers

import (
	"errors"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/docker/machine/libmachine/log"
	"github.com/zhujo01/CorbaTest/libcluster/state"
)

type Driver interface {
	// Create a cluster using the driver's config
	Create() error

	// DriverName returns the name of the driver
	DriverName() string

	// GetCreateFlags returns the mcnflag.Flag slice representing the flags
	// that can be set, their descriptions and defaults.
	GetCreateFlags() []mcnflag.Flag

	// GetClusterName returns the name of the cluster
	GetClusterName() string

	// GetState returns the state that the host is in (running, stopped, etc)
	GetState() (state.State, error)

	// Kill stops a cluster forcefully
	Kill() error

	// PreCreateCheck allows for pre-create operations to make sure a driver is ready for creation
	PreCreateCheck() error

	// Remove a cluster
	Remove() error

	// Restart a cluster
	Restart() error

	// SetConfigFromFlags configures the driver with the object that was returned
	// by RegisterCreateFlags
	SetConfigFromFlags(opts DriverOptions) error

	// Start a cluster
	Start() error

	// Stop a cluster gracefully
	Stop() error

}

var ErrClusterIsNotRunning = errors.New("Cluster is not running")

type DriverOptions interface {
	String(key string) string
	StringSlice(key string) []string
	Int(key string) int
	Bool(key string) bool
}

func ClusterInState(d Driver, desiredState state.State) func() bool {
	return func() bool {
		currentState, err := d.GetState()
		if err != nil {
			log.Debugf("Error getting cluster state: %s", err)
		}
		if currentState == desiredState {
			return true
		}
		return false
	}
}

// MustBeRunning will return an error if the machine is not in a running state.
func MustBeRunning(d Driver) error {
	s, err := d.GetState()
	if err != nil {
		return err
	}

	if s != state.Running {
		return ErrClusterIsNotRunning
	}

	return nil
}








