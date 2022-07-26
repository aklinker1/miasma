package graphql

import (
	"context"
	"strings"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/zero"
)

func validateAppInput(app internal.AppInput) error {
	if app.Name = strings.TrimSpace(app.Name); app.Name == "" {
		return &server.Error{
			Code:    server.EINVALID,
			Message: "App name cannot be empty",
		}
	}
	if app.Image = strings.TrimSpace(app.Image); app.Image == "" {
		return &server.Error{
			Code:    server.EINVALID,
			Message: "App image cannot be empty",
		}
	}
	return nil
}

func (r *Resolver) getAppStatus(ctx context.Context, app internal.App) (internal.RuntimeStatus, error) {
	_, err := r.RuntimeServiceRepo.GetOne(ctx, server.RuntimeServicesFilter{
		AppID: &app.ID,
	})
	if server.ErrorCode(err) == server.ENOTFOUND {
		return internal.RuntimeStatusStopped, nil
	} else if err != nil {
		return internal.RuntimeStatusStopped, err
	}
	return internal.RuntimeStatusRunning, nil
}

func (r *Resolver) deleteAppCascade(ctx context.Context, tx server.Tx, app internal.App) (internal.App, error) {
	deleted, err := r.AppRepo.Delete(ctx, tx, app)
	if err != nil {
		return zero.App, err
	}

	// Routes
	routes, err := r.RouteRepo.GetAll(ctx, tx, server.RoutesFilter{
		AppID: &deleted.ID,
	})
	if err != nil {
		return zero.App, err
	}
	for _, route := range routes {
		_, err = r.RouteRepo.Delete(ctx, tx, route)
		if err != nil {
			return zero.App, err
		}
	}

	// Clear Env
	_, err = r.EnvRepo.Set(ctx, tx, deleted.ID, internal.EnvMap{})
	if err != nil {
		return zero.App, err
	}

	return deleted, nil
}
