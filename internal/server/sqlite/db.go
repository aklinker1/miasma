package sqlite

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"sort"

	_ "github.com/mattn/go-sqlite3"

	"github.com/aklinker1/miasma/internal/server"
)

//go:embed migrations/*.sql
var migrationFS embed.FS

type sqliteDB struct {
	source string
	db     *sql.DB
	logger server.Logger
}

func NewDB(source string, logger server.Logger) server.DB {
	return &sqliteDB{
		source: source,
		logger: logger,
	}
}

func (sqlite *sqliteDB) Open() error {
	sqlite.logger.V("Opening SQLite @ %s", sqlite.source)
	var err error
	sqlite.db, err = sql.Open("sqlite3", sqlite.source)
	if err != nil {
		return err
	}

	// Enable WAL. SQLite performs better with the WAL  because it allows
	// multiple readers to operate while data is being written.
	sqlite.logger.V("Enabling WAL mode in the SQLite database")
	if _, err := sqlite.db.Exec(`PRAGMA journal_mode = wal;`); err != nil {
		return &server.Error{
			Code:            server.EINTERNAL,
			InternalMessage: "Failed to enable wal mode for SQLite",
			ExternalMessage: "Failed to configure SQLite database",
			Op:              "sqliteDB.Open",
			Err:             err,
		}
	}

	ctx := context.Background()
	err = sqlite.migrate(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (sqlite *sqliteDB) migrate(ctx context.Context) error {
	// Create migrations table if needed
	if _, err := sqlite.db.Exec(`CREATE TABLE IF NOT EXISTS migrations (name TEXT PRIMARY KEY);`); err != nil {
		return &server.Error{
			Code:            server.EINTERNAL,
			ExternalMessage: "Failed to create migrations table",
			Op:              "sqliteDB.migrate",
			Err:             err,
		}
	}

	// Get migrations from table
	names, err := fs.Glob(migrationFS, "migrations/*.sql")
	if err != nil {
		return err
	}
	sort.Strings(names)

	// Loop through all SQL migrations, executing each that hasn't been ran
	sqlite.logger.I("Running SQLite migrations...")
	for _, name := range names {
		if err := sqlite.migrateFile(ctx, name); err != nil {
			return fmt.Errorf("migration error: name=%q err=%w", name, err)
		}
	}
	return nil
}

func (sqlite *sqliteDB) migrateFile(ctx context.Context, name string) error {
	tx, err := sqlite.ReadWriteTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Ensure migration has not already been run.
	var n int
	if err := tx.QueryRow(`SELECT COUNT(*) FROM migrations WHERE name = ?`, name).Scan(&n); err != nil {
		return &server.Error{
			Code:            server.EINTERNAL,
			ExternalMessage: fmt.Sprintf("Failed to check if migration '%s' has already been ran", name),
			Op:              "sqliteDB.migrateFile",
			Err:             err,
		}
	} else if n != 0 {
		return nil // already run migration, skip
	}

	// Read and execute migration file.
	sqlite.logger.I(" - %s", name)
	if buf, err := fs.ReadFile(migrationFS, name); err != nil {
		return &server.Error{
			Code:            server.EINTERNAL,
			ExternalMessage: fmt.Sprintf("Failed to read migration file '%s'", name),
			Op:              "sqliteDB.migrateFile",
			Err:             err,
		}
	} else if _, err := tx.Exec(string(buf)); err != nil {
		return &server.Error{
			Code:            server.EINTERNAL,
			ExternalMessage: fmt.Sprintf("Failed to run migration '%s'", name),
			Op:              "sqliteDB.migrateFile",
			Err:             err,
		}
	}

	// Insert record into migrations to prevent re-running migration.
	if _, err := tx.Exec(`INSERT INTO migrations (name) VALUES (?)`, name); err != nil {
		return &server.Error{
			Code:            server.EINTERNAL,
			ExternalMessage: fmt.Sprintf("Failed to save migration '%s' as ran", name),
			Op:              "sqliteDB.migrateFile",
			Err:             err,
		}
	}

	return tx.Commit()
}

func (sqlite *sqliteDB) ReadonlyTx(ctx context.Context) (server.Tx, error) {
	return sqlite.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
}

func (sqlite *sqliteDB) ReadWriteTx(ctx context.Context) (server.Tx, error) {
	return sqlite.db.BeginTx(ctx, nil)
}
