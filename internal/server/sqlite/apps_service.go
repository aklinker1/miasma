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
func (s *AppService) Create(ctx context.Context, app internal.AppInput) (internal.App, error) {
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return EmptyApp, err
	}
	defer tx.Rollback()

	var volumes []internal.BoundVolume
	if len(app.Volumes) > 0 {
		volumes = []internal.BoundVolume{}
		for _, v := range app.Volumes {
			volumes = append(volumes, (internal.BoundVolume)(v))
		}
	}
	created, err := createApp(ctx, tx, internal.App{
		Name:           app.Name,
		Group:          app.Group,
		Image:          app.Image,
		Hidden:         app.Hidden,
		Volumes:        volumes,
		TargetPorts:    app.TargetPorts,
		PublishedPorts: app.PublishedPorts,
		Placement:      app.Placement,
		Networks:       app.Networks,
		Routing:        (*internal.AppRouting)(app.Routing),
		Command:        app.Command,
	})
	if err != nil {
		return EmptyApp, err
	}
	tx.Commit()
	return created, nil
}

// Delete implements server.AppService
func (s *AppService) Delete(ctx context.Context, id string) (internal.App, error) {
	return EmptyApp, server.NewNotImplementedError("sqlite.AppService.Delete")
}

// Get implements server.AppService
func (s *AppService) Get(ctx context.Context, filter internal.AppsFilter) ([]internal.App, error) {
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return nil, err
	}
	return findApps(ctx, tx, filter)
}

// GetOne implements server.AppService
func (s *AppService) GetOne(ctx context.Context, filter internal.AppsFilter) (internal.App, error) {
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
