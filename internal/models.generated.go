// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package internal

import (
	"time"
)

type App struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Group *string `json:"group"`
	// The image and tag the application runs
	Image string `json:"image"`
	// Whether or not the app is returned during regular requests
	Hidden *bool `json:"hidden"`
	// The published ports for the app
	Ports []string `json:"ports"`
	// If the app has routing, a simple string representing that route
	Routing *string `json:"routing"`
	// Whether or not the application is running, stopped, or starting up
	Status string `json:"status"`
	// The number of instances running vs what should be running
	Instances string `json:"instances"`
	// The currently running image digest (hash). Used internally when running
	// applications instead of the tag because the when a new image is pushed, the
	// tag stays the same but the digest changes
	ImageDigest string `json:"imageDigest"`
	// The ports that the app is listening to inside the container. If no target
	// ports are specified, then the container should respect the `PORT` env var.
	TargetPorts []int32 `json:"targetPorts"`
	// The ports that you access the app through in the swarm. This field can, and
	// should be left empty. Miasma automatically manages assigning published ports
	// between 3001-4999. If you need to specify a port, make sure it's outside that
	// range or the port has not been taken. Plugins have set ports starting with
	// 4000, so avoid 4000-4020 if you want to add a plugin at a later date.
	//
	// If these ports are ever cleared, the app will continue using the same ports it
	// was published to before, so that the ports don't change unnecessarily. If you
	// removed it to clear a port for another app/plugin, make sure to restart the
	// app and a new, random port will be allocated for the app, freeing the old
	// port.
	PublishedPorts []int32 `json:"publishedPorts"`
	// The placement constraints specifying which nodes the app will be ran on. Any
	// valid value for the [`--constraint` flag](https://docs.docker.com/engine/swarm/services/#placement-constraints)
	// is valid item in this list
	Placement []string `json:"placement"`
	// Volume bindings for the app
	Volumes []BoundVolume `json:"volumes"`
	// A list of other apps that the service communicates with using their service
	// name and docker's internal DNS. Services don't have to be two way; only the
	// service that accesses the other needs the other network added.
	Networks []string `json:"networks"`
	Command  *string  `json:"command"`
}

type AppChanges struct {
	Name           *string            `json:"name"`
	Image          *string            `json:"image"`
	Group          *string            `json:"group"`
	Hidden         *bool              `json:"hidden"`
	TargetPorts    []int32            `json:"targetPorts"`
	PublishedPorts []int32            `json:"publishedPorts"`
	Placement      []string           `json:"placement"`
	Volumes        []BoundVolumeInput `json:"volumes"`
	Networks       []string           `json:"networks"`
	Command        *string            `json:"command"`
}

type AppInput struct {
	Name           string             `json:"name"`
	Image          string             `json:"image"`
	Group          *string            `json:"group"`
	Hidden         *bool              `json:"hidden"`
	TargetPorts    []int32            `json:"targetPorts"`
	PublishedPorts []int32            `json:"publishedPorts"`
	Placement      []string           `json:"placement"`
	Volumes        []BoundVolumeInput `json:"volumes"`
	Networks       []string           `json:"networks"`
	Command        *string            `json:"command"`
}

type AppRouting struct {
	AppID       string  `json:"appId"`
	Host        *string `json:"host"`
	Path        *string `json:"path"`
	TraefikRule *string `json:"traefikRule"`
}

type AppRoutingInput struct {
	Host        *string `json:"host"`
	Path        *string `json:"path"`
	TraefikRule *string `json:"traefikRule"`
}

type BoundVolume struct {
	// The path inside the container that the data is served from
	Target string `json:"target"`
	// The volume name or directory on the host that the data is stored in
	Source string `json:"source"`
}

type BoundVolumeInput struct {
	Target string `json:"target"`
	Source string `json:"source"`
}

type Health struct {
	// Miasma server's current version
	Version string `json:"version"`
	// The version of docker running on the host, or null if docker is not running
	DockerVersion string `json:"dockerVersion"`
	// The main node's swarm information, or null if not apart of a swarm
	Swarm *SwarmInfo `json:"swarm"`
}

type Plugin struct {
	Name string `json:"name"`
	// Whether or not the plugin has been enabled
	Enable bool `json:"enable"`
}

// The info about the docker swarm if the host running miasma is apart of one.
type SwarmInfo struct {
	ID          string    `json:"id"`
	JoinCommand string    `json:"joinCommand"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}