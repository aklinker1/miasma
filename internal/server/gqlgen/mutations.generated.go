// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgen

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aklinker1/miasma/internal"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type MutationResolver interface {
	CreateApp(ctx context.Context, app internal.AppInput) (*internal.App, error)
	EditApp(ctx context.Context, id string, changes map[string]interface{}) (*internal.App, error)
	DeleteApp(ctx context.Context, id string) (*internal.App, error)
	StartApp(ctx context.Context, id string) (string, error)
	StopApp(ctx context.Context, id string) (string, error)
	ReloadApp(ctx context.Context, id string) (*internal.App, error)
	UpgradeApp(ctx context.Context, id string) (*internal.App, error)
	EnablePlugin(ctx context.Context, pluginName string) (*internal.Plugin, error)
	DisablePlugin(ctx context.Context, pluginName string) (*internal.Plugin, error)
	SetAppRouting(ctx context.Context, appID string, routing *internal.AppRoutingInput) (*internal.AppRouting, error)
	RemoveAppRouting(ctx context.Context, appID string) (*internal.AppRouting, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Mutation_createApp_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 internal.AppInput
	if tmp, ok := rawArgs["app"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("app"))
		arg0, err = ec.unmarshalNAppInput2githubᚗcomᚋaklinker1ᚋmiasmaᚋinternalᚐAppInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["app"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_deleteApp_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
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

func (ec *executionContext) field_Mutation_disablePlugin_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
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

func (ec *executionContext) field_Mutation_editApp_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
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
	var arg1 map[string]interface{}
	if tmp, ok := rawArgs["changes"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("changes"))
		arg1, err = ec.unmarshalNAppChanges2map(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["changes"] = arg1
	return args, nil
}

func (ec *executionContext) field_Mutation_enablePlugin_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
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

func (ec *executionContext) field_Mutation_reloadApp_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
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

func (ec *executionContext) field_Mutation_removeAppRouting_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["appId"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("appId"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["appId"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_setAppRouting_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["appId"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("appId"))
		arg0, err = ec.unmarshalNID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["appId"] = arg0
	var arg1 *internal.AppRoutingInput
	if tmp, ok := rawArgs["routing"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("routing"))
		arg1, err = ec.unmarshalOAppRoutingInput2ᚖgithubᚗcomᚋaklinker1ᚋmiasmaᚋinternalᚐAppRoutingInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["routing"] = arg1
	return args, nil
}

func (ec *executionContext) field_Mutation_startApp_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
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

func (ec *executionContext) field_Mutation_stopApp_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
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

func (ec *executionContext) field_Mutation_upgradeApp_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
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

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Mutation_createApp(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_createApp(ctx, field)
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
		return ec.resolvers.Mutation().CreateApp(rctx, fc.Args["app"].(internal.AppInput))
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

func (ec *executionContext) fieldContext_Mutation_createApp(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
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
			case "group":
				return ec.fieldContext_App_group(ctx, field)
			case "image":
				return ec.fieldContext_App_image(ctx, field)
			case "imageDigest":
				return ec.fieldContext_App_imageDigest(ctx, field)
			case "hidden":
				return ec.fieldContext_App_hidden(ctx, field)
			case "routing":
				return ec.fieldContext_App_routing(ctx, field)
			case "simpleRoute":
				return ec.fieldContext_App_simpleRoute(ctx, field)
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
	if fc.Args, err = ec.field_Mutation_createApp_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_editApp(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_editApp(ctx, field)
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
		return ec.resolvers.Mutation().EditApp(rctx, fc.Args["id"].(string), fc.Args["changes"].(map[string]interface{}))
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

func (ec *executionContext) fieldContext_Mutation_editApp(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
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
			case "group":
				return ec.fieldContext_App_group(ctx, field)
			case "image":
				return ec.fieldContext_App_image(ctx, field)
			case "imageDigest":
				return ec.fieldContext_App_imageDigest(ctx, field)
			case "hidden":
				return ec.fieldContext_App_hidden(ctx, field)
			case "routing":
				return ec.fieldContext_App_routing(ctx, field)
			case "simpleRoute":
				return ec.fieldContext_App_simpleRoute(ctx, field)
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
	if fc.Args, err = ec.field_Mutation_editApp_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_deleteApp(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_deleteApp(ctx, field)
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
		return ec.resolvers.Mutation().DeleteApp(rctx, fc.Args["id"].(string))
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

func (ec *executionContext) fieldContext_Mutation_deleteApp(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
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
			case "group":
				return ec.fieldContext_App_group(ctx, field)
			case "image":
				return ec.fieldContext_App_image(ctx, field)
			case "imageDigest":
				return ec.fieldContext_App_imageDigest(ctx, field)
			case "hidden":
				return ec.fieldContext_App_hidden(ctx, field)
			case "routing":
				return ec.fieldContext_App_routing(ctx, field)
			case "simpleRoute":
				return ec.fieldContext_App_simpleRoute(ctx, field)
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
	if fc.Args, err = ec.field_Mutation_deleteApp_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_startApp(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_startApp(ctx, field)
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
		return ec.resolvers.Mutation().StartApp(rctx, fc.Args["id"].(string))
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
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_startApp(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_startApp_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_stopApp(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_stopApp(ctx, field)
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
		return ec.resolvers.Mutation().StopApp(rctx, fc.Args["id"].(string))
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
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_stopApp(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_stopApp_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_reloadApp(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_reloadApp(ctx, field)
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
		return ec.resolvers.Mutation().ReloadApp(rctx, fc.Args["id"].(string))
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

func (ec *executionContext) fieldContext_Mutation_reloadApp(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
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
			case "group":
				return ec.fieldContext_App_group(ctx, field)
			case "image":
				return ec.fieldContext_App_image(ctx, field)
			case "imageDigest":
				return ec.fieldContext_App_imageDigest(ctx, field)
			case "hidden":
				return ec.fieldContext_App_hidden(ctx, field)
			case "routing":
				return ec.fieldContext_App_routing(ctx, field)
			case "simpleRoute":
				return ec.fieldContext_App_simpleRoute(ctx, field)
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
	if fc.Args, err = ec.field_Mutation_reloadApp_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_upgradeApp(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_upgradeApp(ctx, field)
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
		return ec.resolvers.Mutation().UpgradeApp(rctx, fc.Args["id"].(string))
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

func (ec *executionContext) fieldContext_Mutation_upgradeApp(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
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
			case "group":
				return ec.fieldContext_App_group(ctx, field)
			case "image":
				return ec.fieldContext_App_image(ctx, field)
			case "imageDigest":
				return ec.fieldContext_App_imageDigest(ctx, field)
			case "hidden":
				return ec.fieldContext_App_hidden(ctx, field)
			case "routing":
				return ec.fieldContext_App_routing(ctx, field)
			case "simpleRoute":
				return ec.fieldContext_App_simpleRoute(ctx, field)
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
	if fc.Args, err = ec.field_Mutation_upgradeApp_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_enablePlugin(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_enablePlugin(ctx, field)
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
		return ec.resolvers.Mutation().EnablePlugin(rctx, fc.Args["pluginName"].(string))
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

func (ec *executionContext) fieldContext_Mutation_enablePlugin(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "name":
				return ec.fieldContext_Plugin_name(ctx, field)
			case "enable":
				return ec.fieldContext_Plugin_enable(ctx, field)
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
	if fc.Args, err = ec.field_Mutation_enablePlugin_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_disablePlugin(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_disablePlugin(ctx, field)
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
		return ec.resolvers.Mutation().DisablePlugin(rctx, fc.Args["pluginName"].(string))
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

func (ec *executionContext) fieldContext_Mutation_disablePlugin(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "name":
				return ec.fieldContext_Plugin_name(ctx, field)
			case "enable":
				return ec.fieldContext_Plugin_enable(ctx, field)
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
	if fc.Args, err = ec.field_Mutation_disablePlugin_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_setAppRouting(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_setAppRouting(ctx, field)
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
		return ec.resolvers.Mutation().SetAppRouting(rctx, fc.Args["appId"].(string), fc.Args["routing"].(*internal.AppRoutingInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*internal.AppRouting)
	fc.Result = res
	return ec.marshalOAppRouting2ᚖgithubᚗcomᚋaklinker1ᚋmiasmaᚋinternalᚐAppRouting(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_setAppRouting(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "host":
				return ec.fieldContext_AppRouting_host(ctx, field)
			case "path":
				return ec.fieldContext_AppRouting_path(ctx, field)
			case "traefikRule":
				return ec.fieldContext_AppRouting_traefikRule(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type AppRouting", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_setAppRouting_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_removeAppRouting(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_removeAppRouting(ctx, field)
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
		return ec.resolvers.Mutation().RemoveAppRouting(rctx, fc.Args["appId"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*internal.AppRouting)
	fc.Result = res
	return ec.marshalOAppRouting2ᚖgithubᚗcomᚋaklinker1ᚋmiasmaᚋinternalᚐAppRouting(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_removeAppRouting(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "host":
				return ec.fieldContext_AppRouting_host(ctx, field)
			case "path":
				return ec.fieldContext_AppRouting_path(ctx, field)
			case "traefikRule":
				return ec.fieldContext_AppRouting_traefikRule(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type AppRouting", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_removeAppRouting_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var mutationImplementors = []string{"Mutation"}

func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mutationImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Mutation",
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
			out.Values[i] = graphql.MarshalString("Mutation")
		case "createApp":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_createApp(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "editApp":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_editApp(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "deleteApp":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_deleteApp(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "startApp":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_startApp(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "stopApp":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_stopApp(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "reloadApp":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_reloadApp(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "upgradeApp":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_upgradeApp(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "enablePlugin":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_enablePlugin(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "disablePlugin":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_disablePlugin(ctx, field)
			})

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "setAppRouting":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_setAppRouting(ctx, field)
			})

		case "removeAppRouting":

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_removeAppRouting(ctx, field)
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
