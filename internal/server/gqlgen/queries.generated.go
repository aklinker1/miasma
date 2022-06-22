// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgen

import (
	"context"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/aklinker1/miasma/internal"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type QueryResolver interface {
	Health(ctx context.Context) (*internal.Health, error)
	ListApps(ctx context.Context, page *int32, size *int32, showHidden *bool) ([]*internal.App, error)
	GetApp(ctx context.Context, id string) (*internal.App, error)
	ListPlugins(ctx context.Context) ([]*internal.Plugin, error)
	GetPlugin(ctx context.Context, pluginName string) (*internal.Plugin, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Query___type_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_getApp_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_getPlugin_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["pluginName"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("pluginName"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["pluginName"] = arg0
	return args, nil
}

func (ec *executionContext) field_Query_listApps_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *int32
	if tmp, ok := rawArgs["page"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("page"))
		arg0, err = ec.unmarshalOInt2ᚖint32(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["page"] = arg0
	var arg1 *int32
	if tmp, ok := rawArgs["size"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("size"))
		arg1, err = ec.unmarshalOInt2ᚖint32(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["size"] = arg1
	var arg2 *bool
	if tmp, ok := rawArgs["showHidden"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("showHidden"))
		arg2, err = ec.unmarshalOBoolean2ᚖbool(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["showHidden"] = arg2
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Query_health(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_health(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Health(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*internal.Health)
	fc.Result = res
	return ec.marshalOHealth2ᚖgithubᚗcomᚋaklinker1ᚋmiasmaᚋinternalᚐHealth(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_health(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "version":
				return ec.fieldContext_Health_version(ctx, field)
			case "dockerVersion":
				return ec.fieldContext_Health_dockerVersion(ctx, field)
			case "cluster":
				return ec.fieldContext_Health_cluster(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Health", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Query_listApps(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_listApps(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().ListApps(rctx, fc.Args["page"].(*int32), fc.Args["size"].(*int32), fc.Args["showHidden"].(*bool))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*internal.App)
	fc.Result = res
	return ec.marshalNApp2ᚕᚖgithubᚗcomᚋaklinker1ᚋmiasmaᚋinternalᚐAppᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_listApps(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_App_id(ctx, field)
			case "createdAt":
				return ec.fieldContext_App_createdAt(ctx, field)
			case "updatedAt":
				return ec.fieldContext_App_updatedAt(ctx, field)
			case "name":
				return ec.fieldContext_App_name(ctx, field)
			case "system":
				return ec.fieldContext_App_system(ctx, field)
			case "group":
				return ec.fieldContext_App_group(ctx, field)
			case "image":
				return ec.fieldContext_App_image(ctx, field)
			case "imageDigest":
				return ec.fieldContext_App_imageDigest(ctx, field)
			case "hidden":
				return ec.fieldContext_App_hidden(ctx, field)
			case "route":
				return ec.fieldContext_App_route(ctx, field)
			case "simpleRoute":
				return ec.fieldContext_App_simpleRoute(ctx, field)
			case "env":
				return ec.fieldContext_App_env(ctx, field)
			case "status":
				return ec.fieldContext_App_status(ctx, field)
			case "instances":
				return ec.fieldContext_App_instances(ctx, field)
			case "targetPorts":
				return ec.fieldContext_App_targetPorts(ctx, field)
			case "publishedPorts":
				return ec.fieldContext_App_publishedPorts(ctx, field)
			case "placement":
				return ec.fieldContext_App_placement(ctx, field)
			case "volumes":
				return ec.fieldContext_App_volumes(ctx, field)
			case "networks":
				return ec.fieldContext_App_networks(ctx, field)
			case "command":
				return ec.fieldContext_App_command(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type App", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query_listApps_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Query_getApp(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_getApp(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().GetApp(rctx, fc.Args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*internal.App)
	fc.Result = res
	return ec.marshalNApp2ᚖgithubᚗcomᚋaklinker1ᚋmiasmaᚋinternalᚐApp(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_getApp(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_App_id(ctx, field)
			case "createdAt":
				return ec.fieldContext_App_createdAt(ctx, field)
			case "updatedAt":
				return ec.fieldContext_App_updatedAt(ctx, field)
			case "name":
				return ec.fieldContext_App_name(ctx, field)
			case "system":
				return ec.fieldContext_App_system(ctx, field)
			case "group":
				return ec.fieldContext_App_group(ctx, field)
			case "image":
				return ec.fieldContext_App_image(ctx, field)
			case "imageDigest":
				return ec.fieldContext_App_imageDigest(ctx, field)
			case "hidden":
				return ec.fieldContext_App_hidden(ctx, field)
			case "route":
				return ec.fieldContext_App_route(ctx, field)
			case "simpleRoute":
				return ec.fieldContext_App_simpleRoute(ctx, field)
			case "env":
				return ec.fieldContext_App_env(ctx, field)
			case "status":
				return ec.fieldContext_App_status(ctx, field)
			case "instances":
				return ec.fieldContext_App_instances(ctx, field)
			case "targetPorts":
				return ec.fieldContext_App_targetPorts(ctx, field)
			case "publishedPorts":
				return ec.fieldContext_App_publishedPorts(ctx, field)
			case "placement":
				return ec.fieldContext_App_placement(ctx, field)
			case "volumes":
				return ec.fieldContext_App_volumes(ctx, field)
			case "networks":
				return ec.fieldContext_App_networks(ctx, field)
			case "command":
				return ec.fieldContext_App_command(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type App", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query_getApp_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Query_listPlugins(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_listPlugins(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().ListPlugins(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*internal.Plugin)
	fc.Result = res
	return ec.marshalNPlugin2ᚕᚖgithubᚗcomᚋaklinker1ᚋmiasmaᚋinternalᚐPluginᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_listPlugins(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "name":
				return ec.fieldContext_Plugin_name(ctx, field)
			case "enabled":
				return ec.fieldContext_Plugin_enabled(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Plugin", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Query_getPlugin(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_getPlugin(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().GetPlugin(rctx, fc.Args["pluginName"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*internal.Plugin)
	fc.Result = res
	return ec.marshalNPlugin2ᚖgithubᚗcomᚋaklinker1ᚋmiasmaᚋinternalᚐPlugin(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_getPlugin(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "name":
				return ec.fieldContext_Plugin_name(ctx, field)
			case "enabled":
				return ec.fieldContext_Plugin_enabled(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Plugin", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query_getPlugin_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query___type(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(fc.Args["name"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query___type(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query___type_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query___schema(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema()
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	fc.Result = res
	return ec.marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query___schema(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "description":
				return ec.fieldContext___Schema_description(ctx, field)
			case "types":
				return ec.fieldContext___Schema_types(ctx, field)
			case "queryType":
				return ec.fieldContext___Schema_queryType(ctx, field)
			case "mutationType":
				return ec.fieldContext___Schema_mutationType(ctx, field)
			case "subscriptionType":
				return ec.fieldContext___Schema_subscriptionType(ctx, field)
			case "directives":
				return ec.fieldContext___Schema_directives(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Schema", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var queryImplementors = []string{"Query"}

func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, queryImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Query",
	})

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "health":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_health(ctx, field)
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx, innerFunc)
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return rrm(innerCtx)
			})
		case "listApps":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_listApps(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx, innerFunc)
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return rrm(innerCtx)
			})
		case "getApp":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_getApp(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx, innerFunc)
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return rrm(innerCtx)
			})
		case "listPlugins":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_listPlugins(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx, innerFunc)
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return rrm(innerCtx)
			})
		case "getPlugin":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_getPlugin(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx, innerFunc)
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return rrm(innerCtx)
			})
		case "__type":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Query___type(ctx, field)
			})

		case "__schema":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Query___schema(ctx, field)
			})

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

// endregion ***************************** type.gotpl *****************************
