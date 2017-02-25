package cluster

import (
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/engine"
	"github.com/docker/machine/libmachine/swarm"
	"github.com/docker/machine/libmachine/auth"
)

type Cluster struct {
	ConfigVersion	int
	Driver		drivers.Driver
	DriverName	string
}

type Options struct {
	ClusterDriver	string
	CloudDriver     string
	Size		int
	ImageIds	map[string]string
	EngineOptions 	*engine.Options
	SwarmOptions  	*swarm.Options
	AuthOptions   	*auth.Options
}

type Metadata struct {
	ConfigVersion 	int
	DriverName    	string
	ClusterOptions  Options
}

// cluster can do many things here

