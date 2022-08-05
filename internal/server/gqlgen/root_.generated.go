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
	AppTask() AppTaskResolver
	Health() HealthResolver
	Mutation() MutationResolver
	Node() NodeResolver
	Query() QueryResolver
	Subscription() SubscriptionResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	App struct {
		AutoUpgrade    func(childComplexity int) int
		AvailableAt    func(childComplexity int, clusterIPAddress string) int
		Command        func(childComplexity int) int
		CreatedAt      func(childComplexity int) int
		Env            func(childComplexity int) int
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
		Route          func(childComplexity int) int
		SimpleRoute    func(childComplexity int) int
		Status         func(childComplexity int) int
		System         func(childComplexity int) int
		TargetPorts    func(childComplexity int) int
		UpdatedAt      func(childComplexity int) int
		Volumes        func(childComplexity int) int
	}

	AppInstances struct {
		Running func(childComplexity int) int
		Total   func(childComplexity int) int
	}

	AppTask struct {
		App          func(childComplexity int) int
		AppID        func(childComplexity int) int
		DesiredState func(childComplexity int) int
		Error        func(childComplexity int) int
		ExitCode     func(childComplexity int) int
		Message      func(childComplexity int) int
		Name         func(childComplexity int) int
		Node         func(childComplexity int) int
		NodeID       func(childComplexity int) int
		State        func(childComplexity int) int
		Timestamp    func(childComplexity int) int
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

	Log struct {
		Message   func(childComplexity int) int
		Timestamp func(childComplexity int) int
	}

	Mutation struct {
		CreateApp      func(childComplexity int, input internal.AppInput) int
		DeleteApp      func(childComplexity int, id string) int
		DisablePlugin  func(childComplexity int, name internal.PluginName) int
		EditApp        func(childComplexity int, id string, changes map[string]interface{}) int
		EnablePlugin   func(childComplexity int, name internal.PluginName, config map[string]interface{}) int
		RemoveAppRoute func(childComplexity int, appID string) int
		RestartApp     func(childComplexity int, id string) int
		SetAppEnv      func(childComplexity int, appID string, newEnv map[string]interface{}) int
		SetAppRoute    func(childComplexity int, appID string, route *internal.RouteInput) int
		StartApp       func(childComplexity int, id string) int
		StopApp        func(childComplexity int, id string) int
		UpgradeApp     func(childComplexity int, id string) int
	}

	Node struct {
		Architecture  func(childComplexity int) int
		Hostname      func(childComplexity int) int
		ID            func(childComplexity int) int
		IP            func(childComplexity int) int
		Labels        func(childComplexity int) int
		Os            func(childComplexity int) int
		Services      func(childComplexity int, showHidden *bool) int
		Status        func(childComplexity int) int
		StatusMessage func(childComplexity int) int
	}

	Plugin struct {
		Config  func(childComplexity int) int
		Enabled func(childComplexity int) int
		Name    func(childComplexity int) int
	}

	Query struct {
		GetApp      func(childComplexity int, id string) int
		GetAppTasks func(childComplexity int, id string) int
		GetPlugin   func(childComplexity int, name internal.PluginName) int
		Health      func(childComplexity int) int
		ListApps    func(childComplexity int, page *int, size *int, showHidden *bool) int
		ListPlugins func(childComplexity int) int
		Nodes       func(childComplexity int) int
	}

	Route struct {
		AppID       func(childComplexity int) int
		CreatedAt   func(childComplexity int) int
		Host        func(childComplexity int) int
		Path        func(childComplexity int) int
		TraefikRule func(childComplexity int) int
		UpdatedAt   func(childComplexity int) int
	}

	Subscription struct {
		AppLogs func(childComplexity int, id string, excludeStdout *bool, excludeStderr *bool, initialCount *int) int
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

	case "App.autoUpgrade":
		if e.complexity.App.AutoUpgrade == nil {
			break
		}

		return e.complexity.App.AutoUpgrade(childComplexity), true

	case "App.availableAt":
		if e.complexity.App.AvailableAt == nil {
			break
		}

		args, err := ec.field_App_availableAt_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.App.AvailableAt(childComplexity, args["clusterIpAddress"].(string)), true

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

	case "App.env":
		if e.complexity.App.Env == nil {
			break
		}

		return e.complexity.App.Env(childComplexity), true

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

	case "App.route":
		if e.complexity.App.Route == nil {
			break
		}

		return e.complexity.App.Route(childComplexity), true

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

	case "App.system":
		if e.complexity.App.System == nil {
			break
		}

		return e.complexity.App.System(childComplexity), true

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

	case "AppInstances.running":
		if e.complexity.AppInstances.Running == nil {
			break
		}

		return e.complexity.AppInstances.Running(childComplexity), true

	case "AppInstances.total":
		if e.complexity.AppInstances.Total == nil {
			break
		}

		return e.complexity.AppInstances.Total(childComplexity), true

	case "AppTask.app":
		if e.complexity.AppTask.App == nil {
			break
		}

		return e.complexity.AppTask.App(childComplexity), true

	case "AppTask.appId":
		if e.complexity.AppTask.AppID == nil {
			break
		}

		return e.complexity.AppTask.AppID(childComplexity), true

	case "AppTask.desiredState":
		if e.complexity.AppTask.DesiredState == nil {
			break
		}

		return e.complexity.AppTask.DesiredState(childComplexity), true

	case "AppTask.error":
		if e.complexity.AppTask.Error == nil {
			break
		}

		return e.complexity.AppTask.Error(childComplexity), true

	case "AppTask.exitCode":
		if e.complexity.AppTask.ExitCode == nil {
			break
		}

		return e.complexity.AppTask.ExitCode(childComplexity), true

	case "AppTask.message":
		if e.complexity.AppTask.Message == nil {
			break
		}

		return e.complexity.AppTask.Message(childComplexity), true

	case "AppTask.name":
		if e.complexity.AppTask.Name == nil {
			break
		}

		return e.complexity.AppTask.Name(childComplexity), true

	case "AppTask.node":
		if e.complexity.AppTask.Node == nil {
			break
		}

		return e.complexity.AppTask.Node(childComplexity), true

	case "AppTask.nodeId":
		if e.complexity.AppTask.NodeID == nil {
			break
		}

		return e.complexity.AppTask.NodeID(childComplexity), true

	case "AppTask.state":
		if e.complexity.AppTask.State == nil {
			break
		}

		return e.complexity.AppTask.State(childComplexity), true

	case "AppTask.timestamp":
		if e.complexity.AppTask.Timestamp == nil {
			break
		}

		return e.complexity.AppTask.Timestamp(childComplexity), true

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

	case "Log.message":
		if e.complexity.Log.Message == nil {
			break
		}

		return e.complexity.Log.Message(childComplexity), true

	case "Log.timestamp":
		if e.complexity.Log.Timestamp == nil {
			break
		}

		return e.complexity.Log.Timestamp(childComplexity), true

	case "Mutation.createApp":
		if e.complexity.Mutation.CreateApp == nil {
			break
		}

		args, err := ec.field_Mutation_createApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateApp(childComplexity, args["input"].(internal.AppInput)), true

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

		return e.complexity.Mutation.DisablePlugin(childComplexity, args["name"].(internal.PluginName)), true

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

		return e.complexity.Mutation.EnablePlugin(childComplexity, args["name"].(internal.PluginName), args["config"].(map[string]interface{})), true

	case "Mutation.removeAppRoute":
		if e.complexity.Mutation.RemoveAppRoute == nil {
			break
		}

		args, err := ec.field_Mutation_removeAppRoute_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RemoveAppRoute(childComplexity, args["appId"].(string)), true

	case "Mutation.restartApp":
		if e.complexity.Mutation.RestartApp == nil {
			break
		}

		args, err := ec.field_Mutation_restartApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RestartApp(childComplexity, args["id"].(string)), true

	case "Mutation.setAppEnv":
		if e.complexity.Mutation.SetAppEnv == nil {
			break
		}

		args, err := ec.field_Mutation_setAppEnv_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SetAppEnv(childComplexity, args["appId"].(string), args["newEnv"].(map[string]interface{})), true

	case "Mutation.setAppRoute":
		if e.complexity.Mutation.SetAppRoute == nil {
			break
		}

		args, err := ec.field_Mutation_setAppRoute_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SetAppRoute(childComplexity, args["appId"].(string), args["route"].(*internal.RouteInput)), true

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

	case "Node.architecture":
		if e.complexity.Node.Architecture == nil {
			break
		}

		return e.complexity.Node.Architecture(childComplexity), true

	case "Node.hostname":
		if e.complexity.Node.Hostname == nil {
			break
		}

		return e.complexity.Node.Hostname(childComplexity), true

	case "Node.id":
		if e.complexity.Node.ID == nil {
			break
		}

		return e.complexity.Node.ID(childComplexity), true

	case "Node.ip":
		if e.complexity.Node.IP == nil {
			break
		}

		return e.complexity.Node.IP(childComplexity), true

	case "Node.labels":
		if e.complexity.Node.Labels == nil {
			break
		}

		return e.complexity.Node.Labels(childComplexity), true

	case "Node.os":
		if e.complexity.Node.Os == nil {
			break
		}

		return e.complexity.Node.Os(childComplexity), true

	case "Node.services":
		if e.complexity.Node.Services == nil {
			break
		}

		args, err := ec.field_Node_services_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Node.Services(childComplexity, args["showHidden"].(*bool)), true

	case "Node.status":
		if e.complexity.Node.Status == nil {
			break
		}

		return e.complexity.Node.Status(childComplexity), true

	case "Node.statusMessage":
		if e.complexity.Node.StatusMessage == nil {
			break
		}

		return e.complexity.Node.StatusMessage(childComplexity), true

	case "Plugin.config":
		if e.complexity.Plugin.Config == nil {
			break
		}

		return e.complexity.Plugin.Config(childComplexity), true

	case "Plugin.enabled":
		if e.complexity.Plugin.Enabled == nil {
			break
		}

		return e.complexity.Plugin.Enabled(childComplexity), true

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

	case "Query.getAppTasks":
		if e.complexity.Query.GetAppTasks == nil {
			break
		}

		args, err := ec.field_Query_getAppTasks_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetAppTasks(childComplexity, args["id"].(string)), true

	case "Query.getPlugin":
		if e.complexity.Query.GetPlugin == nil {
			break
		}

		args, err := ec.field_Query_getPlugin_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetPlugin(childComplexity, args["name"].(internal.PluginName)), true

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

		return e.complexity.Query.ListApps(childComplexity, args["page"].(*int), args["size"].(*int), args["showHidden"].(*bool)), true

	case "Query.listPlugins":
		if e.complexity.Query.ListPlugins == nil {
			break
		}

		return e.complexity.Query.ListPlugins(childComplexity), true

	case "Query.nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		return e.complexity.Query.Nodes(childComplexity), true

	case "Route.appId":
		if e.complexity.Route.AppID == nil {
			break
		}

		return e.complexity.Route.AppID(childComplexity), true

	case "Route.createdAt":
		if e.complexity.Route.CreatedAt == nil {
			break
		}

		return e.complexity.Route.CreatedAt(childComplexity), true

	case "Route.host":
		if e.complexity.Route.Host == nil {
			break
		}

		return e.complexity.Route.Host(childComplexity), true

	case "Route.path":
		if e.complexity.Route.Path == nil {
			break
		}

		return e.complexity.Route.Path(childComplexity), true

	case "Route.traefikRule":
		if e.complexity.Route.TraefikRule == nil {
			break
		}

		return e.complexity.Route.TraefikRule(childComplexity), true

	case "Route.updatedAt":
		if e.complexity.Route.UpdatedAt == nil {
			break
		}

		return e.complexity.Route.UpdatedAt(childComplexity), true

	case "Subscription.appLogs":
		if e.complexity.Subscription.AppLogs == nil {
			break
		}

		args, err := ec.field_Subscription_appLogs_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Subscription.AppLogs(childComplexity, args["id"].(string), args["excludeStdout"].(*bool), args["excludeStderr"].(*bool), args["initialCount"].(*int)), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputAppInput,
		ec.unmarshalInputBoundVolumeInput,
		ec.unmarshalInputRouteInput,
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
	case ast.Subscription:
		next := ec._Subscription(ctx, rc.Operation.SelectionSet)

		var buf bytes.Buffer
		return func(ctx context.Context) *graphql.Response {
			buf.Reset()
			data := next(ctx)

			if data == nil {
				return nil
			}
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
	{Name: "../../../api/models.graphqls", Input: `"""
Contains useful information about the cluster.
"""
type ClusterInfo {
  "The Docker swarm ID"
  id: String!
  "The command to run on other machines to join the cluster"
  joinCommand: String!
  "When the cluster was initialized"
  createdAt: Time!
  "When the cluster was last updated"
  updatedAt: Time!
}

"Server health and version information"
type Health {
  "Miasma server's current version."
  version: String!
  "The version of docker running on the host, or null if docker is not running."
  dockerVersion: String!
  "The cluster versioning and information, or ` + "`" + `null` + "`" + ` if not apart of a cluster."
  cluster: ClusterInfo
}

"Docker volume configuration"
type BoundVolume {
  "The path inside the container that the data is served from."
  target: String!
  "The volume name or directory on the host that the data is stored in."
  source: String!
}

enum RuntimeStatus {
  RUNNING
  STOPPED
}

"Managed application"
type App {
  id: ID!
  createdAt: Time!
  updatedAt: Time!
  "The app name. Different from the docker service name, which is the name but lower case and all spaces replaced with dashes"
  name: String!
  "Whether or not the application is managed by the system. You cannot edit or delete system apps."
  system: Boolean!
  "A string used to group the app"
  group: String
  "The image and tag the application runs."
  image: String!
  """
  The currently running image digest (hash). Used internally when running
  applications instead of the tag because the when a new image is pushed, the
  tag stays the same but the digest changes.
  """
  imageDigest: String!
  """
  Whether or not the app should automatically upgrade when a newer version of it's image is available. Defaults to ` + "`" + `true` + "`" + ` when creating an app

  App upgrades are automatically checked according the the ` + "`" + `AUTO_UPDATE_CRON` + "`" + ` expression.
  """
  autoUpgrade: Boolean!
  "Whether or not the app is returned during regular requests."
  hidden: Boolean!
  "If the app has a route and the traefik plugin is enabled, this is it's config."
  route: Route
  "If the app has a route and the traefik plugin is enabled, this is a simple representation of it."
  simpleRoute: String
  """
  A list of URLs the application can be accessed at, including the ` + "`" + `simpleRoute` + "`" + `, and all the published ports
  """
  availableAt(clusterIpAddress: String!): [String!]!
  "The environment variables configured for this app."
  env: Map
  "Whether or not the application is running, or stopped."
  status: RuntimeStatus!
  "The number of instances running vs what should be running."
  instances: AppInstances!
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
  is valid item in this list.
  """
  placement: [String!]
  "Volume bindings for the app."
  volumes: [BoundVolume!]
  """
  A list of other apps that the service communicates with using their service
  name and docker's internal DNS. Services don't have to be two way; only the
  service that accesses the other needs the other network added.
  """
  networks: [String!]
  "Custom docker command. This is an array of arguments starting with the binary that is being executed"
  command: [String!]
}

"Input type for [BoundVolume](#boundvolume)."
input BoundVolumeInput {
  target: String!
  source: String!
}

"Input type for [App](#app)."
input AppInput {
  name: String!
  image: String!
  autoUpgrade: Boolean
  group: String
  hidden: Boolean
  targetPorts: [Int!]
  publishedPorts: [Int!]
  placement: [String!]
  volumes: [BoundVolumeInput!]
  networks: [String!]
  command: [String!]
}

"[Changeset](#Changeset) input type for [App](#app)."
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
  command: [String!]
}

"Plugins are apps with deeper integrations with Miasma."
type Plugin {
  name: PluginName!
  "Whether or not the plugin has been enabled."
  enabled: Boolean!
  "Plugin's configuration."
  config: Map!
}

"Rules around where an app can be accessed from."
type Route {
  appId: ID!
  createdAt: Time!
  updatedAt: Time!
  "The URL's hostname, ex: 'example.com' or 'google.com'."
  host: String
  "A custom path at the end of the URL: ex: '/search' or '/console'"
  path: String
  """
  A custom Traefik rule instead of just a host and path, ex: '(Host(domain1.com) || Host(domain2.com)'

  See [Traefik's docs](https://doc.traefik.io/traefik/routing/routers/#rule) for usage and complex examples.
  """
  traefikRule: String
}

"Input type for [Route](#route)."
input RouteInput {
  host: String
  path: String
  traefikRule: String
}

"Contains information about how many instances of the app are running vs supposed to be running"
type AppInstances {
  running: Int!
  total: Int!
}

"Unique identifier for plugins"
enum PluginName {
  "The name of the [Traefik](https://doc.traefik.io/traefik/) ingress router plugin"
  TRAEFIK
}

"Details about a machine in the cluster."
type Node {
  "The docker node's ID."
  id: String!
  "The OS the node is running"
  os: String!
  "The CPU architecture of the node. Services are automatically placed on nodes based on their image's supported architectures and the nodes' architectures."
  architecture: String!
  "The machines hostname, as returned by the ` + "`" + `hostname` + "`" + ` command."
  hostname: String!
  "The IP address the node joined the cluster as."
  ip: String!

  "` + "`" + `unknown` + "`" + `, ` + "`" + `down` + "`" + `, ` + "`" + `ready` + "`" + `, or ` + "`" + `disconnected` + "`" + `. See Docker's [API docs](https://docs.docker.com/engine/api/v1.41/#operation/NodeInspect)."
  status: String!
  "The node's status message, usually present when when the status is not ` + "`" + `ready` + "`" + `."
  statusMessage: String
  "The node's labels, mostly used to place apps on specific nodes."
  labels: Map!
  "List of apps running on the machine"
  services(
    "Same as ` + "`" + `listApps` + "`" + `'s argument. When ` + "`" + `true` + "`" + `, hidden apps will be returned"
    showHidden: Boolean = false
  ): [App!]!
}

"Tasks define the desired state of on app. If you're familiar with docker, this returns the result of ` + "`" + `docker service ps` + "`" + `"
type AppTask {
  message: String!
  state: String!
  desiredState: String!
  timestamp: Time!
  appId: String!
  app: App!
  nodeId: String!
  node: Node!
  name: String!
  error: String
  exitCode: Int
}

type Log {
  message: String!
  timestamp: Time!
}
`, BuiltIn: false},
	{Name: "../../../api/mutations.graphqls", Input: `type Mutation {
  "Create and start a new app."
  createApp(input: AppInput!): App!
  "Edit app configuration."
  editApp(id: ID!, changes: AppChanges!): App!
  "Stop and delete an app."
  deleteApp(id: ID!): App!
  "Start a stopped app."
  startApp(id: ID!): App!
  "Stop a running app."
  stopApp(id: ID!): App!
  "Stop and restart an app."
  restartApp(id: ID!): App!
  "Pull the latest version of the app's image and then restart."
  upgradeApp(id: ID!): App!

  "Enable one of Miasma's plugins."
  enablePlugin(
    "The name of the plugin to enable."
    name: PluginName!
    "Any plugin specific configuration."
    config: Map
  ): Plugin!
  "Disable one of Miasma's plugins."
  disablePlugin("The name of the plugin to disable." name: PluginName!): Plugin!

  "Set an app's environnement variables"
  setAppEnv(
    appId: ID!
    "A map of variable names to their values. Docker only supports UPPER_SNAKE_CASE variable names"
    newEnv: Map
  ): Map

  """
  Set's an app's route.

  Only available when the 'router' plugin is enabled
  """
  setAppRoute(appId: ID!, route: RouteInput): Route
  """
  Removes an app's route.

  Only available when the 'router' plugin is enabled
  """
  removeAppRoute(appId: ID!): Route
}
`, BuiltIn: false},
	{Name: "../../../api/queries.graphqls", Input: `type Query {
  "Get the server's health and version information"
  health: Health

  "List the running apps"
  listApps(
    "The page to start at for pagination, the first page is 1."
    page: Int = 1
    "Number of apps to return per page."
    size: Int = 10
    "Whether or not to includes apps that are marked as hidden."
    showHidden: Boolean = false
  ): [App!]!
  "Grab an app by it's ID"
  getApp(id: ID!): App!

  "List of tasks for an app, up the 5 most recent"
  getAppTasks(id: ID!): [AppTask!]!

  "List all the available plugins for Miasma"
  listPlugins: [Plugin!]!
  "Grab a plugin by it's name"
  getPlugin(name: PluginName!): Plugin!

  "List the nodes that are apart of the cluster"
  nodes: [Node!]!
}
`, BuiltIn: false},
	{Name: "../../../api/scalars.graphqls", Input: `"A JSON map of key-value pairs. Values can be any type."
scalar Map

"ISO 8601 date time string."
scalar Time
`, BuiltIn: false},
	{Name: "../../../api/subscription.graphqls", Input: `type Subscription {
  "Returns the latest log one at a time."
  appLogs(
    "The ID of the app you want to listen for logs from."
    id: ID!
    excludeStdout: Boolean = false
    excludeStderr: Boolean = false
    "The subscription will load this number of logs from the past initially before listening for future logs."
    initialCount: Int = 50
  ): Log!
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
