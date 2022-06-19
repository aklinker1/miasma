package sqlite

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
)

var (
	EmptyApp = internal.App{}
)

type AppService struct {
	db server.DB
}

func NewAppService(db server.DB) server.AppService {
	return &AppService{
		db: db,
	}
}

// Create implements server.AppService
func (s *AppService) Create(ctx context.Context, app internal.App) (internal.App, error) {
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return EmptyApp, err
	}
	defer tx.Rollback()

	created, err := createApp(ctx, tx, app)
	if err != nil {
		return EmptyApp, err
	}

	if app.Routing != nil {
		createdRoute, err := createRoute(ctx, tx, internal.AppRouting{
			AppID:       created.ID,
			Host:        app.Routing.Host,
			Path:        app.Routing.Path,
			TraefikRule: app.Routing.TraefikRule,
		})
		if err != nil {
			return EmptyApp, err
		}
		created.Routing = &createdRoute
	}

	tx.Commit()
	return created, nil
}

// Delete implements server.AppService
func (s *AppService) Delete(ctx context.Context, id string) (internal.App, error) {
	tx, err := s.db.ReadWriteTx(ctx)
	if err != nil {
		return EmptyApp, err
	}
	defer tx.Rollback()

	app, err := findApp(ctx, tx, server.AppsFilter{
		ID: &id,
	})
	if err != nil {
		return EmptyApp, err
	}
	route, err := findRoute(ctx, tx, server.RoutesFilter{
		AppID: &id,
	})
	if err != nil {
		return EmptyApp, err
	}
	app.Routing = &route

	err = deleteApp(ctx, tx, app)
	if err != nil {
		return EmptyApp, err
	}
	err = deleteRoute(ctx, tx, route)
	if err != nil {
		return EmptyApp, err
	}

	tx.Commit()
	return app, nil
}

func (s *AppService) FindApps(ctx context.Context, filter server.AppsFilter) ([]internal.App, error) {
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return nil, err
	}
	return findApps(ctx, tx, filter)
}

// GetOne implements server.AppService
func (s *AppService) FindApp(ctx context.Context, filter server.AppsFilter) (internal.App, error) {
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return EmptyApp, err
	}
	return findApp(ctx, tx, filter)
}

// Update implements server.AppService
func (s *AppService) Update(ctx context.Context, app internal.App) (internal.App, error) {
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return EmptyApp, err
	}
	defer tx.Rollback()

	created, err := updateApp(ctx, tx, app)
	if err != nil {
		return EmptyApp, err
	}
	tx.Commit()
	return created, nil
}
