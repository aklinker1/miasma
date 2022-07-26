package zero

import (
	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/docker/docker/api/types/swarm"
)

var (
	App                = internal.App{}
	Plugin             = internal.Plugin{}
	Route              = internal.Route{}
	SwarmService       = swarm.Service{}
	SwarmNode          = swarm.Node{}
	SwarmTask          = swarm.Task{}
	AppInstances       = internal.AppInstances{}
	RuntimeServiceSpec = server.RuntimeServiceSpec{}
)
