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
	db      server.DB
	runtime server.RuntimeService
	logger  server.Logger
}

func NewAppService(db server.DB, runtime server.RuntimeService, logger server.Logger) server.AppService {
	return &AppService{
		db:      db,
		runtime: runtime,
		logger:  logger,
	}
}

// Create implements server.AppService
func (s *AppService) Create(ctx context.Context, app internal.App) (internal.App, error) {
	s.logger.D("Creating app: %s", app.Name)
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return EmptyApp, err
	}
	defer tx.Rollback()

	// Get image digest
	imageDigest, err := s.runtime.PullLatest(ctx, app.Image)
	if err != nil {
		return EmptyApp, err
	}
	app.ImageDigest = imageDigest

	// Create app
	created, err := createApp(ctx, tx, app)
	if err != nil {
		return EmptyApp, err
	}

	// Create routing if necessary
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

	// Start the app
	err = s.runtime.Start(ctx, created)
	if err != nil {
		return EmptyApp, err
	}

	tx.Commit()
	return created, nil
}

// Delete implements server.AppService
func (s *AppService) Delete(ctx context.Context, id string) (internal.App, error) {
	s.logger.D("Creating app: %s", id)
	tx, err := s.db.ReadWriteTx(ctx)
	if err != nil {
		return EmptyApp, err
	}
	defer tx.Rollback()

	// Find DB models
	app, err := findApp(ctx, tx, server.AppsFilter{
		ID: &id,
	})
	if err != nil {
		return EmptyApp, err
	}
	route, err := findRoute(ctx, tx, server.RoutesFilter{
		AppID: &id,
	})
	if server.ErrorCode(err) == server.ENOTFOUND {
		// noop
	} else if err != nil {
		return EmptyApp, err
	} else {
		app.Routing = &route
	}

	// Remove DB models
	err = deleteApp(ctx, tx, app)
	if err != nil {
		return EmptyApp, err
	}
	err = deleteRoute(ctx, tx, route)
	if err != nil {
		return EmptyApp, err
	}

	// After all other successes, stop the service
	err = s.runtime.Stop(ctx, app)
	if err != nil {
		return EmptyApp, err
	}

	tx.Commit()
	return app, nil
}

func (s *AppService) FindApps(ctx context.Context, filter server.AppsFilter) ([]internal.App, error) {
	s.logger.D("Finding apps that match: %+v", filter)
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return nil, err
	}
	return findApps(ctx, tx, filter)
}

// GetOne implements server.AppService
func (s *AppService) FindApp(ctx context.Context, filter server.AppsFilter) (internal.App, error) {
	s.logger.D("Finding an app that matches: %+v", filter)
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return EmptyApp, err
	}
	return findApp(ctx, tx, filter)
}

// Update implements server.AppService
func (s *AppService) Update(ctx context.Context, app internal.App, newImage *string) (internal.App, error) {
	s.logger.D("Updating app: %s", app.Name)
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return EmptyApp, err
	}
	defer tx.Rollback()

	if app.System {
		return EmptyApp, &server.Error{
			Code:    server.EINVALID,
			Message: "Cannot edit Miasma's system apps",
			Op:      "sqlite.AppService.Update",
		}
	}

	if newImage != nil {
		newDigest, err := s.runtime.PullLatest(ctx, *newImage)
		if err != nil {
			return EmptyApp, err
		}
		app.ImageDigest = newDigest
	}

	created, err := updateApp(ctx, tx, app)
	if err != nil {
		return EmptyApp, err
	}

	err = s.runtime.Restart(ctx, app)
	if err != nil {
		return EmptyApp, err
	}

	tx.Commit()
	return created, nil
}
