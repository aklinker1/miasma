package sqlite

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
)

var (
	EmptyEnvMap      = internal.EnvMap{}
	EmptyEnvVariable = envVariable{}
)

type EnvService struct {
	db      server.DB
	runtime server.RuntimeService
	logger  server.Logger
}

func NewEnvService(db server.DB, runtime server.RuntimeService, logger server.Logger) server.EnvService {
	return &EnvService{
		db:      db,
		logger:  logger,
		runtime: runtime,
	}
}

// FindEnv implements server.EnvService
func (s *EnvService) FindEnv(ctx context.Context, filter server.EnvFilter) (internal.EnvMap, error) {
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return EmptyEnvMap, err
	}
	defer tx.Rollback()

	env, err := findEnvMap(ctx, tx, filter)
	if err != nil {
		return EmptyEnvMap, err
	}

	tx.Commit()
	return env, nil
}

// SetAppEnv implements server.EnvService
func (s *EnvService) SetAppEnv(ctx context.Context, appID string, newEnv internal.EnvMap) (internal.EnvMap, error) {
	tx, err := s.db.ReadWriteTx(ctx)
	if err != nil {
		return EmptyEnvMap, err
	}
	defer tx.Rollback()

	// Delete and recreate the env
	err = deleteAppEnv(ctx, tx, appID)
	if err != nil {
		return EmptyEnvMap, err
	}
	for key, value := range newEnv {
		_, err := createEnvVariable(ctx, tx, envVariable{
			AppID: appID,
			Key:   key,
			Value: value,
		})
		if err != nil {
			return EmptyEnvMap, err
		}
	}

	// Restart the app
	app, route, _, plugins, err := findStartParams(ctx, tx, appID, startParamKnowns{
		env:      newEnv,
		knownEnv: true,
	})
	if err != nil {
		return EmptyEnvMap, err
	}
	err = s.runtime.Restart(ctx, app, route, newEnv, plugins)
	if err != nil {
		return EmptyEnvMap, err
	}

	tx.Commit()
	return newEnv, nil
}
