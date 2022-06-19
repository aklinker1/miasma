// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgen

import (
	"bytes"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/aklinker1/miasma/internal"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	App() AppResolver
	Health() HealthResolver
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	App struct {
		Command        func(childComplexity int) int
		CreatedAt      func(childComplexity int) int
		Group          func(childComplexity int) int
		Hidden         func(childComplexity int) int
		ID             func(childComplexity int) int
		Image          func(childComplexity int) int
		ImageDigest    func(childComplexity int) int
		Instances      func(childComplexity int) int
		Name           func(childComplexity int) int
		Networks       func(childComplexity int) int
		Placement      func(childComplexity int) int
		PublishedPorts func(childComplexity int) int
		Routing        func(childComplexity int) int
		SimpleRoute    func(childComplexity int) int
		Status         func(childComplexity int) int
		TargetPorts    func(childComplexity int) int
		UpdatedAt      func(childComplexity int) int
		Volumes        func(childComplexity int) int
	}

	AppRouting struct {
		Host        func(childComplexity int) int
		Path        func(childComplexity int) int
		TraefikRule func(childComplexity int) int
	}

	BoundVolume struct {
		Source func(childComplexity int) int
		Target func(childComplexity int) int
	}

	ClusterInfo struct {
		CreatedAt   func(childComplexity int) int
		ID          func(childComplexity int) int
		JoinCommand func(childComplexity int) int
		UpdatedAt   func(childComplexity int) int
	}

	Health struct {
		Cluster       func(childComplexity int) int
		DockerVersion func(childComplexity int) int
		Version       func(childComplexity int) int
	}

	Mutation struct {
		CreateApp        func(childComplexity int, app internal.AppInput) int
		DeleteApp        func(childComplexity int, id string) int
		DisablePlugin    func(childComplexity int, pluginName string) int
		EditApp          func(childComplexity int, id string, changes map[string]interface{}) int
		EnablePlugin     func(childComplexity int, pluginName string) int
		ReloadApp        func(childComplexity int, id string) int
		RemoveAppRouting func(childComplexity int, appID string) int
		SetAppRouting    func(childComplexity int, appID string, routing *internal.AppRoutingInput) int
		StartApp         func(childComplexity int, id string) int
		StopApp          func(childComplexity int, id string) int
		UpgradeApp       func(childComplexity int, id string) int
	}

	Plugin struct {
		Enable func(childComplexity int) int
		Name   func(childComplexity int) int
	}

	Query struct {
		GetApp        func(childComplexity int, id string) int
		GetAppRouting func(childComplexity int, appID string) int
		GetPlugin     func(childComplexity int, pluginName string) int
		Health        func(childComplexity int) int
		ListApps      func(childComplexity int, page *int32, size *int32, showHidden *bool) int
		ListPlugins   func(childComplexity int) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "App.command":
		if e.complexity.App.Command == nil {
			break
		}

		return e.complexity.App.Command(childComplexity), true

	case "App.createdAt":
		if e.complexity.App.CreatedAt == nil {
			break
		}

		return e.complexity.App.CreatedAt(childComplexity), true

	case "App.group":
		if e.complexity.App.Group == nil {
			break
		}

		return e.complexity.App.Group(childComplexity), true

	case "App.hidden":
		if e.complexity.App.Hidden == nil {
			break
		}

		return e.complexity.App.Hidden(childComplexity), true

	case "App.id":
		if e.complexity.App.ID == nil {
			break
		}

		return e.complexity.App.ID(childComplexity), true

	case "App.image":
		if e.complexity.App.Image == nil {
			break
		}

		return e.complexity.App.Image(childComplexity), true

	case "App.imageDigest":
		if e.complexity.App.ImageDigest == nil {
			break
		}

		return e.complexity.App.ImageDigest(childComplexity), true

	case "App.instances":
		if e.complexity.App.Instances == nil {
			break
		}

		return e.complexity.App.Instances(childComplexity), true

	case "App.name":
		if e.complexity.App.Name == nil {
			break
		}

		return e.complexity.App.Name(childComplexity), true

	case "App.networks":
		if e.complexity.App.Networks == nil {
			break
		}

		return e.complexity.App.Networks(childComplexity), true

	case "App.placement":
		if e.complexity.App.Placement == nil {
			break
		}

		return e.complexity.App.Placement(childComplexity), true

	case "App.publishedPorts":
		if e.complexity.App.PublishedPorts == nil {
			break
		}

		return e.complexity.App.PublishedPorts(childComplexity), true

	case "App.routing":
		if e.complexity.App.Routing == nil {
			break
		}

		return e.complexity.App.Routing(childComplexity), true

	case "App.simpleRoute":
		if e.complexity.App.SimpleRoute == nil {
			break
		}

		return e.complexity.App.SimpleRoute(childComplexity), true

	case "App.status":
		if e.complexity.App.Status == nil {
			break
		}

		return e.complexity.App.Status(childComplexity), true

	case "App.targetPorts":
		if e.complexity.App.TargetPorts == nil {
			break
		}

		return e.complexity.App.TargetPorts(childComplexity), true

	case "App.updatedAt":
		if e.complexity.App.UpdatedAt == nil {
			break
		}

		return e.complexity.App.UpdatedAt(childComplexity), true

	case "App.volumes":
		if e.complexity.App.Volumes == nil {
			break
		}

		return e.complexity.App.Volumes(childComplexity), true

	case "AppRouting.host":
		if e.complexity.AppRouting.Host == nil {
			break
		}

		return e.complexity.AppRouting.Host(childComplexity), true

	case "AppRouting.path":
		if e.complexity.AppRouting.Path == nil {
			break
		}

		return e.complexity.AppRouting.Path(childComplexity), true

	case "AppRouting.traefikRule":
		if e.complexity.AppRouting.TraefikRule == nil {
			break
		}

		return e.complexity.AppRouting.TraefikRule(childComplexity), true

	case "BoundVolume.source":
		if e.complexity.BoundVolume.Source == nil {
			break
		}

		return e.complexity.BoundVolume.Source(childComplexity), true

	case "BoundVolume.target":
		if e.complexity.BoundVolume.Target == nil {
			break
		}

		return e.complexity.BoundVolume.Target(childComplexity), true

	case "ClusterInfo.createdAt":
		if e.complexity.ClusterInfo.CreatedAt == nil {
			break
		}

		return e.complexity.ClusterInfo.CreatedAt(childComplexity), true

	case "ClusterInfo.id":
		if e.complexity.ClusterInfo.ID == nil {
			break
		}

		return e.complexity.ClusterInfo.ID(childComplexity), true

	case "ClusterInfo.joinCommand":
		if e.complexity.ClusterInfo.JoinCommand == nil {
			break
		}

		return e.complexity.ClusterInfo.JoinCommand(childComplexity), true

	case "ClusterInfo.updatedAt":
		if e.complexity.ClusterInfo.UpdatedAt == nil {
			break
		}

		return e.complexity.ClusterInfo.UpdatedAt(childComplexity), true

	case "Health.cluster":
		if e.complexity.Health.Cluster == nil {
			break
		}

		return e.complexity.Health.Cluster(childComplexity), true

	case "Health.dockerVersion":
		if e.complexity.Health.DockerVersion == nil {
			break
		}

		return e.complexity.Health.DockerVersion(childComplexity), true

	case "Health.version":
		if e.complexity.Health.Version == nil {
			break
		}

		return e.complexity.Health.Version(childComplexity), true

	case "Mutation.createApp":
		if e.complexity.Mutation.CreateApp == nil {
			break
		}

		args, err := ec.field_Mutation_createApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateApp(childComplexity, args["app"].(internal.AppInput)), true

	case "Mutation.deleteApp":
		if e.complexity.Mutation.DeleteApp == nil {
			break
		}

		args, err := ec.field_Mutation_deleteApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteApp(childComplexity, args["id"].(string)), true

	case "Mutation.disablePlugin":
		if e.complexity.Mutation.DisablePlugin == nil {
			break
		}

		args, err := ec.field_Mutation_disablePlugin_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DisablePlugin(childComplexity, args["pluginName"].(string)), true

	case "Mutation.editApp":
		if e.complexity.Mutation.EditApp == nil {
			break
		}

		args, err := ec.field_Mutation_editApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.EditApp(childComplexity, args["id"].(string), args["changes"].(map[string]interface{})), true

	case "Mutation.enablePlugin":
		if e.complexity.Mutation.EnablePlugin == nil {
			break
		}

		args, err := ec.field_Mutation_enablePlugin_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.EnablePlugin(childComplexity, args["pluginName"].(string)), true

	case "Mutation.reloadApp":
		if e.complexity.Mutation.ReloadApp == nil {
			break
		}

		args, err := ec.field_Mutation_reloadApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.ReloadApp(childComplexity, args["id"].(string)), true

	case "Mutation.removeAppRouting":
		if e.complexity.Mutation.RemoveAppRouting == nil {
			break
		}

		args, err := ec.field_Mutation_removeAppRouting_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RemoveAppRouting(childComplexity, args["appId"].(string)), true

	case "Mutation.setAppRouting":
		if e.complexity.Mutation.SetAppRouting == nil {
			break
		}

		args, err := ec.field_Mutation_setAppRouting_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SetAppRouting(childComplexity, args["appId"].(string), args["routing"].(*internal.AppRoutingInput)), true

	case "Mutation.startApp":
		if e.complexity.Mutation.StartApp == nil {
			break
		}

		args, err := ec.field_Mutation_startApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.StartApp(childComplexity, args["id"].(string)), true

	case "Mutation.stopApp":
		if e.complexity.Mutation.StopApp == nil {
			break
		}

		args, err := ec.field_Mutation_stopApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.StopApp(childComplexity, args["id"].(string)), true

	case "Mutation.upgradeApp":
		if e.complexity.Mutation.UpgradeApp == nil {
			break
		}

		args, err := ec.field_Mutation_upgradeApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpgradeApp(childComplexity, args["id"].(string)), true

	case "Plugin.enable":
		if e.complexity.Plugin.Enable == nil {
			break
		}

		return e.complexity.Plugin.Enable(childComplexity), true

	case "Plugin.name":
		if e.complexity.Plugin.Name == nil {
			break
		}

		return e.complexity.Plugin.Name(childComplexity), true

	case "Query.getApp":
		if e.complexity.Query.GetApp == nil {
			break
		}

		args, err := ec.field_Query_getApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetApp(childComplexity, args["id"].(string)), true

	case "Query.getAppRouting":
		if e.complexity.Query.GetAppRouting == nil {
			break
		}

		args, err := ec.field_Query_getAppRouting_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetAppRouting(childComplexity, args["appId"].(string)), true

	case "Query.getPlugin":
		if e.complexity.Query.GetPlugin == nil {
			break
		}

		args, err := ec.field_Query_getPlugin_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetPlugin(childComplexity, args["pluginName"].(string)), true

	case "Query.health":
		if e.complexity.Query.Health == nil {
			break
		}

		return e.complexity.Query.Health(childComplexity), true

	case "Query.listApps":
		if e.complexity.Query.ListApps == nil {
			break
		}

		args, err := ec.field_Query_listApps_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.ListApps(childComplexity, args["page"].(*int32), args["size"].(*int32), args["showHidden"].(*bool)), true

	case "Query.listPlugins":
		if e.complexity.Query.ListPlugins == nil {
			break
		}

		return e.complexity.Query.ListPlugins(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputAppInput,
		ec.unmarshalInputAppRoutingInput,
		ec.unmarshalInputBoundVolumeInput,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "api/models.graphqls", Input: `"""
The info about the docker swarm if the host running miasma is apart of one.
"""
type ClusterInfo {
  id: String!
  joinCommand: String!
  createdAt: Time!
  updatedAt: Time!
}

type Health {
  "Miasma server's current version"
  version: String!
  "The version of docker running on the host, or null if docker is not running"
  dockerVersion: String!
  "The cluster versioning and information, or ` + "`" + `null` + "`" + ` if not apart of a cluster"
  cluster: ClusterInfo
}

type BoundVolume {
  "The path inside the container that the data is served from"
  target: String!
  "The volume name or directory on the host that the data is stored in"
  source: String!
}

type App {
  id: ID!
  createdAt: Time!
  updatedAt: Time!
  name: String!
  group: String
  "The image and tag the application runs"
  image: String!
  """
  The currently running image digest (hash). Used internally when running
  applications instead of the tag because the when a new image is pushed, the
  tag stays the same but the digest changes
  """
  imageDigest: String!
  "Whether or not the app is returned during regular requests"
  hidden: Boolean!
  "If the app has routing, this is the routing config"
  routing: AppRouting
  "If the app has routing, a simple string representing that route"
  simpleRoute: String
  "Whether or not the application is running, stopped, or starting up"
  status: String!
  "The number of instances running vs what should be running"
  instances: String!
  """
  The ports that the app is listening to inside the container. If no target
  ports are specified, then the container should respect the ` + "`" + `PORT` + "`" + ` env var.
  """
  targetPorts: [Int!]
  """
  The ports that you access the app through in the swarm. This field can, and
  should be left empty. Miasma automatically manages assigning published ports
  between 3001-4999. If you need to specify a port, make sure it's outside that
  range or the port has not been taken. Plugins have set ports starting with
  4000, so avoid 4000-4020 if you want to add a plugin at a later date.

  If these ports are ever cleared, the app will continue using the same ports it
  was published to before, so that the ports don't change unnecessarily. If you
  removed it to clear a port for another app/plugin, make sure to restart the
  app and a new, random port will be allocated for the app, freeing the old
  port.
  """
  publishedPorts: [Int!]
  """
  The placement constraints specifying which nodes the app will be ran on. Any
  valid value for the [` + "`" + `--constraint` + "`" + ` flag](https://docs.docker.com/engine/swarm/services/#placement-constraints)
  is valid item in this list
  """
  placement: [String!]
  "Volume bindings for the app"
  volumes: [BoundVolume!]
  """
  A list of other apps that the service communicates with using their service
  name and docker's internal DNS. Services don't have to be two way; only the
  service that accesses the other needs the other network added.
  """
  networks: [String!]
  command: String
}

input BoundVolumeInput {
  target: String!
  source: String!
}

input AppInput {
  name: String!
  image: String!
  group: String
  hidden: Boolean
  targetPorts: [Int!]
  publishedPorts: [Int!]
  placement: [String!]
  volumes: [BoundVolumeInput!]
  networks: [String!]
  routing: AppRoutingInput
  command: String
}

input AppChanges {
  name: String
  image: String
  group: String
  hidden: Boolean
  targetPorts: [Int!]
  publishedPorts: [Int!]
  placement: [String!]
  volumes: [BoundVolumeInput!]
  networks: [String!]
  command: String
}

type Plugin {
  name: String!
  "Whether or not the plugin has been enabled"
  enable: Boolean!
}

type AppRouting {
  host: String
  path: String
  traefikRule: String
}

input AppRoutingInput {
  host: String
  path: String
  traefikRule: String
}
`, BuiltIn: false},
	{Name: "api/mutations.graphqls", Input: `type Mutation {
  "Create and start a new app"
  createApp(app: AppInput!): App!
  "Edit app metadata unrelated to how the container(s) that are run"
  editApp(id: ID!, changes: AppChanges!): App!
  "Stop and delete an app"
  deleteApp(id: ID!): App!
  "Start a stopped app"
  startApp(id: ID!): String!
  "Stop a running app"
  stopApp(id: ID!): String!
  "Stop and restart an app"
  reloadApp(id: ID!): App!
  "Pull the latest version of the app's image and then restart"
  upgradeApp(id: ID!): App!

  "Install one of Miasma's plugins"
  enablePlugin(pluginName: String!): Plugin!
  "Disable one of Miasma's plugins"
  disablePlugin(pluginName: String!): Plugin!

  "Only available when the 'router' plugin is enabled"
  setAppRouting(appId: ID!, routing: AppRoutingInput): AppRouting
  "Only available when the 'router' plugin is enabled"
  removeAppRouting(appId: ID!): AppRouting
}
`, BuiltIn: false},
	{Name: "api/queries.graphqls", Input: `type Query {
  health: Health

  listApps(page: Int = 1, size: Int = 10, showHidden: Boolean): [App!]!
  getApp(id: ID!): App!

  listPlugins: [Plugin!]!
  getPlugin(pluginName: String!): Plugin!

  "Only available when the 'router' plugin is enabled"
  getAppRouting(appId: ID!): AppRouting!
}
`, BuiltIn: false},
	{Name: "api/scalars.graphqls", Input: `scalar Map
scalar Time
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
