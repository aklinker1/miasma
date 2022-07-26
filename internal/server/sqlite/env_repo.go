package sqlite

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/sqlite/sqlb"
	"github.com/aklinker1/miasma/internal/server/zero"
	"github.com/samber/lo"
)

type envVariable struct {
	AppID string
	Key   string
	Value string
}

type EnvRepo struct {
	Logger server.Logger
}

func (r *EnvRepo) Get(ctx context.Context, tx server.Tx, filter server.EnvFilter) (internal.EnvMap, error) {
	env, err := r.getAllEnvVars(ctx, tx, filter)
	if err != nil {
		return zero.EnvMap, err
	}
	return lo.Reduce(env, func(res internal.EnvMap, v envVariable, _ int) internal.EnvMap {
		res[v.Key] = v.Value
		return res
	}, internal.EnvMap{}), nil
}

func (r *EnvRepo) Set(ctx context.Context, tx server.Tx, appID string, newEnv internal.EnvMap) (internal.EnvMap, error) {
	// Delete current env vars
	filter := server.EnvFilter{AppID: appID}
	err := r.deleteAll(ctx, tx, filter)
	if err != nil {
		return zero.EnvMap, err
	}

	// Create new ones
	for key, value := range newEnv {
		_, err = r.createEnvVar(ctx, tx, envVariable{
			AppID: appID,
			Key:   key,
			Value: value,
		})
		if err != nil {
			return nil, err
		}
	}
	return newEnv, nil
}

func (r *EnvRepo) getAllEnvVars(ctx context.Context, tx server.Tx, filter server.EnvFilter) ([]envVariable, error) {
	var scanned envVariable
	query := sqlb.Select(r.Logger, "env", map[string]any{
		"app_id": &scanned.AppID,
		"key":    &scanned.Key,
		"value":  &scanned.Value,
	}).Where("app_id = ?", filter.AppID)

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, server.NewDatabaseError("getAllEnvVars", err)
	}
	dest := query.ScanDest()
	result := make([]envVariable, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, server.NewDatabaseError("getAllEnvVars", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
}

func (r *EnvRepo) createEnvVar(ctx context.Context, tx server.Tx, envVar envVariable) (envVariable, error) {
	sql, args := sqlb.Insert(r.Logger, "env", map[string]any{
		"app_id": envVar.AppID,
		"key":    envVar.Key,
		"value":  envVar.Value,
	}).ToSQL()
	_, err := tx.ExecContext(ctx, sql, args...)
	return envVar, err
}

func (r *EnvRepo) deleteAll(ctx context.Context, tx server.Tx, filter server.EnvFilter) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM env WHERE app_id = ?", filter.AppID)
	return err
}
