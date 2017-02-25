package kubernetesazure


import (
	"github.com/zhujo01/CorbaTest/libcluster/drivers"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/zhujo01/CorbaTest/libcluster/state"
)

// Kubernetes on Azure (implementation of cluster driver interface)



type Driver struct {
	*drivers.BaseDriver
}

func (d *Driver) Create() error {
	return nil
}

func (d *Driver) DriverName() string {
	return "kubernetesazure"
}

func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	return nil
}

func (d *Driver) GetClusterName() string {
	return nil
}

func (d *Driver) GetState() (state.State, error) {
	return nil
}

func (d *Driver) Kill() error {
	return nil
}

func (d *Driver) PreCreateCheck() error {
	return nil
}

func (d *Driver)  Remove() error {
	return nil
}

func (d *Driver) Restart() error {
	return nil
}

func (d *Driver) SetConfigFromFlags(opts drivers.DriverOptions) error {
	return nil
}


func (d *Driver) Start() error {
	return nil
}

func (d *Driver) Stop() error {
	return nil
}
