// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package internal

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

// Managed application
type App struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// The app name. Different from the docker service name, which is the name but lower case and all spaces replaced with dashes
	Name string `json:"name"`
	// Whether or not the application is managed by the system. You cannot edit or delete system apps.
	System bool `json:"system"`
	// A string used to group the app
	Group *string `json:"group"`
	// The image and tag the application runs.
	Image string `json:"image"`
	// The currently running image digest (hash). Used internally when running
	// applications instead of the tag because the when a new image is pushed, the
	// tag stays the same but the digest changes.
	ImageDigest string `json:"imageDigest"`
	// Whether or not the app should automatically upgrade when a newer version of it's image is available. Defaults to `true` when creating an app
	//
	// App upgrades are automatically checked according the the `AUTO_UPDATE_CRON` expression.
	AutoUpgrade bool `json:"autoUpgrade"`
	// Whether or not the app is returned during regular requests.
	Hidden bool `json:"hidden"`
	// If the app has a route and the traefik plugin is enabled, this is it's config.
	Route *Route `json:"route"`
	// If the app has a route and the traefik plugin is enabled, this is a simple representation of it.
	SimpleRoute *string `json:"simpleRoute"`
	// A list of URLs the application can be accessed at, including the `simpleRoute`, and all the published ports
	AvailableAt []string `json:"availableAt"`
	// The environment variables configured for this app.
	Env map[string]interface{} `json:"env"`
	// Whether or not the application is running, or stopped.
	Status string `json:"status"`
	// The number of instances running vs what should be running.
	Instances *AppInstances `json:"instances"`
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
	// is valid item in this list.
	Placement []string `json:"placement"`
	// Volume bindings for the app.
	Volumes []*BoundVolume `json:"volumes"`
	// A list of other apps that the service communicates with using their service
	// name and docker's internal DNS. Services don't have to be two way; only the
	// service that accesses the other needs the other network added.
	Networks []string `json:"networks"`
	// Custom docker command. This is an array of arguments starting with the binary that is being executed
	Command []string `json:"command"`
}

// Input type for [App](#app).
type AppInput struct {
	Name           string              `json:"name"`
	Image          string              `json:"image"`
	AutoUpgrade    *bool               `json:"autoUpgrade"`
	Group          *string             `json:"group"`
	Hidden         *bool               `json:"hidden"`
	TargetPorts    []int32             `json:"targetPorts"`
	PublishedPorts []int32             `json:"publishedPorts"`
	Placement      []string            `json:"placement"`
	Volumes        []*BoundVolumeInput `json:"volumes"`
	Networks       []string            `json:"networks"`
	Command        []string            `json:"command"`
}

// Contains information about how many instances of the app are running vs supposed to be running
type AppInstances struct {
	Running int32 `json:"running"`
	Total   int32 `json:"total"`
}

// Docker volume configuration
type BoundVolume struct {
	// The path inside the container that the data is served from.
	Target string `json:"target"`
	// The volume name or directory on the host that the data is stored in.
	Source string `json:"source"`
}

// Input type for [BoundVolume](#boundvolume).
type BoundVolumeInput struct {
	Target string `json:"target"`
	Source string `json:"source"`
}

// Contains useful information about the cluster.
type ClusterInfo struct {
	// The Docker swarm ID
	ID string `json:"id"`
	// The command to run on other machines to join the cluster
	JoinCommand string `json:"joinCommand"`
	// When the cluster was initialized
	CreatedAt time.Time `json:"createdAt"`
	// When the cluster was last updated
	UpdatedAt time.Time `json:"updatedAt"`
}

// Server health and version information
type Health struct {
	// Miasma server's current version.
	Version string `json:"version"`
	// The version of docker running on the host, or null if docker is not running.
	DockerVersion string `json:"dockerVersion"`
	// The cluster versioning and information, or `null` if not apart of a cluster.
	Cluster *ClusterInfo `json:"cluster"`
}

type Log struct {
	Message string `json:"message"`
}

// Details about a machine in the cluster.
type Node struct {
	// The docker node's ID.
	ID string `json:"id"`
	// The OS the node is running
	Os string `json:"os"`
	// The CPU architecture of the node. Services are automatically placed on nodes based on their image's supported architectures and the nodes' architectures.
	Architecture string `json:"architecture"`
	// The machines hostname, as returned by the `hostname` command.
	Hostname string `json:"hostname"`
	// The IP address the node joined the cluster as.
	IP string `json:"ip"`
	// `unknown`, `down`, `ready`, or `disconnected`. See Docker's [API docs](https://docs.docker.com/engine/api/v1.41/#operation/NodeInspect).
	Status string `json:"status"`
	// The node's status message, usually present when when the status is not `ready`.
	StatusMessage *string `json:"statusMessage"`
	// The node's labels, mostly used to place apps on specific nodes.
	Labels map[string]interface{} `json:"labels"`
	// List of apps running on the machine
	Services []*App `json:"services"`
}

// Plugins are apps with deeper integrations with Miasma.
type Plugin struct {
	Name PluginName `json:"name"`
	// Whether or not the plugin has been enabled.
	Enabled bool `json:"enabled"`
	// Plugin's configuration.
	Config map[string]interface{} `json:"config"`
}

// Rules around where an app can be accessed from.
type Route struct {
	AppID     string    `json:"appId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// The URL's hostname, ex: 'example.com' or 'google.com'.
	Host *string `json:"host"`
	// A custom path at the end of the URL: ex: '/search' or '/console'
	Path *string `json:"path"`
	// A custom Traefik rule instead of just a host and path, ex: '(Host(domain1.com) || Host(domain2.com)'
	//
	// See [Traefik's docs](https://doc.traefik.io/traefik/routing/routers/#rule) for usage and complex examples.
	TraefikRule *string `json:"traefikRule"`
}

// Input type for [Route](#route).
type RouteInput struct {
	Host        *string `json:"host"`
	Path        *string `json:"path"`
	TraefikRule *string `json:"traefikRule"`
}

// Unique identifier for plugins
type PluginName string

const (
	// The name of the [Traefik](https://doc.traefik.io/traefik/) ingress router plugin
	PluginNameTraefik PluginName = "TRAEFIK"
)

var AllPluginName = []PluginName{
	PluginNameTraefik,
}

func (e PluginName) IsValid() bool {
	switch e {
	case PluginNameTraefik:
		return true
	}
	return false
}

func (e PluginName) String() string {
	return string(e)
}

func (e *PluginName) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PluginName(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PluginName", str)
	}
	return nil
}

func (e PluginName) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
