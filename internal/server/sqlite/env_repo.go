package sqlite

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/sqlite/sqlb"
	"github.com/samber/lo"
)

type envVariable struct {
	AppID string
	Key   string
	Value string
}

func findEnv(ctx context.Context, tx server.Tx, filter server.EnvFilter) ([]envVariable, error) {
	var scanned envVariable
	query := sqlb.Select("env", map[string]any{
		"app_id": &scanned.AppID,
		"key":    &scanned.Key,
		"value":  &scanned.Value,
	})
	if filter.AppID != nil {
		query.Where("app_id = ?", *filter.AppID)
	}

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, server.NewDatabaseError("findEnv", err)
	}
	dest := query.ScanDest()
	result := make([]envVariable, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, server.NewDatabaseError("findEnv", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
}

func findEnvMap(ctx context.Context, tx server.Tx, filter server.EnvFilter) (internal.EnvMap, error) {
	env, err := findEnv(ctx, tx, filter)
	if err != nil {
		return EmptyEnvMap, err
	}
	return lo.Reduce(env, func(res internal.EnvMap, v envVariable, _ int) internal.EnvMap {
		res[v.Key] = v.Value
		return res
	}, internal.EnvMap{}), nil
}

func createEnvVariable(ctx context.Context, tx server.Tx, envVar envVariable) (envVariable, error) {
	sql, args := sqlb.Insert("env", map[string]any{
		"app_id": envVar.AppID,
		"key":    envVar.Key,
		"value":  envVar.Value,
	}).ToSQL()
	_, err := tx.ExecContext(ctx, sql, args...)
	return envVar, err
}

func deleteAppEnv(ctx context.Context, tx server.Tx, appID string) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM env WHERE app_id = ?", appID)
	return err
}
