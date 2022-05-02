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
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	App struct {
		Group     func(childComplexity int) int
		Hidden    func(childComplexity int) int
		ID        func(childComplexity int) int
		Image     func(childComplexity int) int
		Instances func(childComplexity int) int
		Name      func(childComplexity int) int
		Ports     func(childComplexity int) int
		Routing   func(childComplexity int) int
		Status    func(childComplexity int) int
	}

	AppRouting struct {
		AppID       func(childComplexity int) int
		Host        func(childComplexity int) int
		Path        func(childComplexity int) int
		TraefikRule func(childComplexity int) int
	}

	BoundVolume struct {
		Source func(childComplexity int) int
		Target func(childComplexity int) int
	}

	DockerConfig struct {
		AppID          func(childComplexity int) int
		Command        func(childComplexity int) int
		ImageDigest    func(childComplexity int) int
		Networks       func(childComplexity int) int
		Placement      func(childComplexity int) int
		PublishedPorts func(childComplexity int) int
		TargetPorts    func(childComplexity int) int
		Volumes        func(childComplexity int) int
	}

	Group struct {
		Apps func(childComplexity int) int
		Name func(childComplexity int) int
	}

	Health struct {
		DockerVersion func(childComplexity int) int
		Swarm         func(childComplexity int) int
		Version       func(childComplexity int) int
	}

	Mutation struct {
		CreateApp          func(childComplexity int, app internal.CreateAppInput) int
		DeleteApp          func(childComplexity int, appName string) int
		DisablePlugin      func(childComplexity int, pluginName string) int
		EditApp            func(childComplexity int, appName string, app internal.EditAppInput) int
		EnablePlugin       func(childComplexity int, pluginName string) int
		ReloadApp          func(childComplexity int, appName string) int
		RemoveAppRouting   func(childComplexity int, appName string) int
		SetAppDockerConfig func(childComplexity int, appName string, newConfig *internal.DockerConfigInput) int
		SetAppEnv          func(childComplexity int, appName string, newEnv map[string]interface{}) int
		SetAppRouting      func(childComplexity int, appName string, routing *internal.AppRoutingInput) int
		StartApp           func(childComplexity int, appName string) int
		StopApp            func(childComplexity int, appName string) int
		UpgradeApp         func(childComplexity int, appName string) int
	}

	Plugin struct {
		Enable func(childComplexity int) int
		Name   func(childComplexity int) int
	}

	Query struct {
		GetApp             func(childComplexity int, appName string) int
		GetAppDockerConfig func(childComplexity int, appName string) int
		GetAppEnv          func(childComplexity int, appName string) int
		GetAppRouting      func(childComplexity int, appName string) int
		GetPlugin          func(childComplexity int, pluginName string) int
		Health             func(childComplexity int) int
		ListApps           func(childComplexity int, page *int32, size *int32, showHidden *bool) int
		ListPlugins        func(childComplexity int) int
	}

	SwarmInfo struct {
		CreatedAt   func(childComplexity int) int
		ID          func(childComplexity int) int
		JoinCommand func(childComplexity int) int
		UpdatedAt   func(childComplexity int) int
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

	case "App.ports":
		if e.complexity.App.Ports == nil {
			break
		}

		return e.complexity.App.Ports(childComplexity), true

	case "App.routing":
		if e.complexity.App.Routing == nil {
			break
		}

		return e.complexity.App.Routing(childComplexity), true

	case "App.status":
		if e.complexity.App.Status == nil {
			break
		}

		return e.complexity.App.Status(childComplexity), true

	case "AppRouting.appId":
		if e.complexity.AppRouting.AppID == nil {
			break
		}

		return e.complexity.AppRouting.AppID(childComplexity), true

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

	case "DockerConfig.appId":
		if e.complexity.DockerConfig.AppID == nil {
			break
		}

		return e.complexity.DockerConfig.AppID(childComplexity), true

	case "DockerConfig.command":
		if e.complexity.DockerConfig.Command == nil {
			break
		}

		return e.complexity.DockerConfig.Command(childComplexity), true

	case "DockerConfig.imageDigest":
		if e.complexity.DockerConfig.ImageDigest == nil {
			break
		}

		return e.complexity.DockerConfig.ImageDigest(childComplexity), true

	case "DockerConfig.networks":
		if e.complexity.DockerConfig.Networks == nil {
			break
		}

		return e.complexity.DockerConfig.Networks(childComplexity), true

	case "DockerConfig.placement":
		if e.complexity.DockerConfig.Placement == nil {
			break
		}

		return e.complexity.DockerConfig.Placement(childComplexity), true

	case "DockerConfig.publishedPorts":
		if e.complexity.DockerConfig.PublishedPorts == nil {
			break
		}

		return e.complexity.DockerConfig.PublishedPorts(childComplexity), true

	case "DockerConfig.targetPorts":
		if e.complexity.DockerConfig.TargetPorts == nil {
			break
		}

		return e.complexity.DockerConfig.TargetPorts(childComplexity), true

	case "DockerConfig.volumes":
		if e.complexity.DockerConfig.Volumes == nil {
			break
		}

		return e.complexity.DockerConfig.Volumes(childComplexity), true

	case "Group.apps":
		if e.complexity.Group.Apps == nil {
			break
		}

		return e.complexity.Group.Apps(childComplexity), true

	case "Group.name":
		if e.complexity.Group.Name == nil {
			break
		}

		return e.complexity.Group.Name(childComplexity), true

	case "Health.dockerVersion":
		if e.complexity.Health.DockerVersion == nil {
			break
		}

		return e.complexity.Health.DockerVersion(childComplexity), true

	case "Health.swarm":
		if e.complexity.Health.Swarm == nil {
			break
		}

		return e.complexity.Health.Swarm(childComplexity), true

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

		return e.complexity.Mutation.CreateApp(childComplexity, args["app"].(internal.CreateAppInput)), true

	case "Mutation.deleteApp":
		if e.complexity.Mutation.DeleteApp == nil {
			break
		}

		args, err := ec.field_Mutation_deleteApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteApp(childComplexity, args["appName"].(string)), true

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

		return e.complexity.Mutation.EditApp(childComplexity, args["appName"].(string), args["app"].(internal.EditAppInput)), true

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

		return e.complexity.Mutation.ReloadApp(childComplexity, args["appName"].(string)), true

	case "Mutation.removeAppRouting":
		if e.complexity.Mutation.RemoveAppRouting == nil {
			break
		}

		args, err := ec.field_Mutation_removeAppRouting_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RemoveAppRouting(childComplexity, args["appName"].(string)), true

	case "Mutation.setAppDockerConfig":
		if e.complexity.Mutation.SetAppDockerConfig == nil {
			break
		}

		args, err := ec.field_Mutation_setAppDockerConfig_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SetAppDockerConfig(childComplexity, args["appName"].(string), args["newConfig"].(*internal.DockerConfigInput)), true

	case "Mutation.setAppEnv":
		if e.complexity.Mutation.SetAppEnv == nil {
			break
		}

		args, err := ec.field_Mutation_setAppEnv_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SetAppEnv(childComplexity, args["appName"].(string), args["newEnv"].(map[string]interface{})), true

	case "Mutation.setAppRouting":
		if e.complexity.Mutation.SetAppRouting == nil {
			break
		}

		args, err := ec.field_Mutation_setAppRouting_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SetAppRouting(childComplexity, args["appName"].(string), args["routing"].(*internal.AppRoutingInput)), true

	case "Mutation.startApp":
		if e.complexity.Mutation.StartApp == nil {
			break
		}

		args, err := ec.field_Mutation_startApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.StartApp(childComplexity, args["appName"].(string)), true

	case "Mutation.stopApp":
		if e.complexity.Mutation.StopApp == nil {
			break
		}

		args, err := ec.field_Mutation_stopApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.StopApp(childComplexity, args["appName"].(string)), true

	case "Mutation.upgradeApp":
		if e.complexity.Mutation.UpgradeApp == nil {
			break
		}

		args, err := ec.field_Mutation_upgradeApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpgradeApp(childComplexity, args["appName"].(string)), true

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

		return e.complexity.Query.GetApp(childComplexity, args["appName"].(string)), true

	case "Query.getAppDockerConfig":
		if e.complexity.Query.GetAppDockerConfig == nil {
			break
		}

		args, err := ec.field_Query_getAppDockerConfig_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetAppDockerConfig(childComplexity, args["appName"].(string)), true

	case "Query.getAppEnv":
		if e.complexity.Query.GetAppEnv == nil {
			break
		}

		args, err := ec.field_Query_getAppEnv_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetAppEnv(childComplexity, args["appName"].(string)), true

	case "Query.getAppRouting":
		if e.complexity.Query.GetAppRouting == nil {
			break
		}

		args, err := ec.field_Query_getAppRouting_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetAppRouting(childComplexity, args["appName"].(string)), true

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

	case "SwarmInfo.createdAt":
		if e.complexity.SwarmInfo.CreatedAt == nil {
			break
		}

		return e.complexity.SwarmInfo.CreatedAt(childComplexity), true

	case "SwarmInfo.id":
		if e.complexity.SwarmInfo.ID == nil {
			break
		}

		return e.complexity.SwarmInfo.ID(childComplexity), true

	case "SwarmInfo.joinCommand":
		if e.complexity.SwarmInfo.JoinCommand == nil {
			break
		}

		return e.complexity.SwarmInfo.JoinCommand(childComplexity), true

	case "SwarmInfo.updatedAt":
		if e.complexity.SwarmInfo.UpdatedAt == nil {
			break
		}

		return e.complexity.SwarmInfo.UpdatedAt(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputAppRoutingInput,
		ec.unmarshalInputBoundVolumeInput,
		ec.unmarshalInputCreateAppInput,
		ec.unmarshalInputDockerConfigInput,
		ec.unmarshalInputEditAppInput,
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
type SwarmInfo {
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
  "The main node's swarm information, or null if not apart of a swarm"
  swarm: SwarmInfo
}

type Group {
  "A simple label to track what apps are related"
  name: String!
  "The apps in the group"
  apps: [App!]!
}

type App {
  id: ID!
  name: String!
  "The group the app belongs to, or ` + "`" + `null` + "`" + ` if it doesn't belong to a group"
  group: Group
  "The image and tag the application runs"
  image: String!
  "Whether or not the app is returned during regular requests"
  hidden: Boolean
  "The published ports for the app"
  ports: [String!]
  "If the app has routing, a simple string representing that route"
  routing: String
  "Whether or not the application is running, stopped, or starting up"
  status: String!
  "The number of instances running vs what should be running"
  instances: String!
}

input CreateAppInput {
  name: String!
  image: String!
  groupName: String
  hidden: Boolean
}

input EditAppInput {
  name: String!
  groupName: String
  hidden: Boolean
}

type BoundVolume {
  "The path inside the container that the data is served from"
  target: String!
  "The volume name or directory on the host that the data is stored in"
  source: String!
}

input BoundVolumeInput {
  target: String!
  source: String!
}

type DockerConfig {
  "The ID of the app the run config is for"
  appId: ID!
  """
  The currently running image digest (hash). Used internally when running
  applications instead of the tag because the when a new image is pushed, the
  tag stays the same but the digest changes
  """
  imageDigest: String!
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

"Excluded inputs will be considered as empty inputs, clearing each empty field"
input DockerConfigInput {
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
  appId: ID!
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
  createApp(app: CreateAppInput!): App!
  "Edit app metadata unrelated to how the container(s) that are run"
  editApp(appName: String!, app: EditAppInput!): App!
  "Stop and delete an app"
  deleteApp(appName: String!): App!
  "Start a stopped app"
  startApp(appName: String!): String!
  "Stop a running app"
  stopApp(appName: String!): String!
  "Stop and restart an app"
  reloadApp(appName: String!): App!
  "Pull the latest version of the app's image and then restart"
  upgradeApp(appName: String!): App!

  "Configure the docker runtime config of an app then reload the app"
  setAppDockerConfig(appName: String!, newConfig: DockerConfigInput): DockerConfig!

  """
  Update the app's environment variables all at once. Excluded environment
  variables are removed
  """
  setAppEnv(appName: String!, newEnv: Map!): Map!

  "Install one of Miasma's plugins"
  enablePlugin(pluginName: String!): Plugin!
  "Disable one of Miasma's plugins"
  disablePlugin(pluginName: String!): Plugin!

  "Only available when the 'router' plugin is enabled"
  setAppRouting(appName: String!, routing: AppRoutingInput): AppRouting
  "Only available when the 'router' plugin is enabled"
  removeAppRouting(appName: String!): AppRouting
}
`, BuiltIn: false},
	{Name: "api/queries.graphqls", Input: `type Query {
  health: Health

  listApps(page: Int, size: Int, showHidden: Boolean): [App!]!
  getApp(appName: String!): App!

  getAppDockerConfig(appName: String!): DockerConfig!

  getAppEnv(appName: String!): Map!

  listPlugins: [Plugin!]!
  getPlugin(pluginName: String!): Plugin!

  "Only available when the 'router' plugin is enabled"
  getAppRouting(appName: String!): AppRouting!
}
`, BuiltIn: false},
	{Name: "api/scalars.graphqls", Input: `scalar Map
scalar Time
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
