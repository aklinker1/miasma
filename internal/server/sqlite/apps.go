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
	return EmptyApp, server.NewNotImplementedError("sqlite.AppService.Create")
}

// Delete implements server.AppService
func (s *AppService) Delete(ctx context.Context, appName string) (internal.App, error) {
	return EmptyApp, server.NewNotImplementedError("sqlite.AppService.Delete")
}

// Get implements server.AppService
func (s *AppService) Get(ctx context.Context, options server.GetAppOptions) ([]internal.App, error) {
	return nil, server.NewNotImplementedError("sqlite.AppService.Get")
}

// Update implements server.AppService
func (s *AppService) Update(ctx context.Context, app internal.App) error {
	return server.NewNotImplementedError("sqlite.AppService.Update")
}
