package sqlite

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
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
	panic("unimplemented")
}

// Delete implements server.AppService
func (s *AppService) Delete(ctx context.Context, appName string) (internal.App, error) {
	panic("unimplemented")
}

// Get implements server.AppService
func (s *AppService) Get(ctx context.Context, options server.GetAppOptions) ([]internal.App, error) {
	panic("unimplemented")
}

// Update implements server.AppService
func (s *AppService) Update(ctx context.Context, app internal.App) error {
	panic("unimplemented")
}
