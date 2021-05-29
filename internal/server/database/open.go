package database

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/database/entities"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func open() *gorm.DB {
	if db != nil {
		return db
	}

	var err error
	db, err = gorm.Open(sqlite.Open("/data/miasma/apps.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println(err)
		panic("Failed to open SQLite database")
	}

	err = db.AutoMigrate(
		&models.App{},
		&entities.SQLRunConfig{},
		&entities.SQLEnvVar{},
		&models.Plugin{},
		&models.TraefikPluginConfig{},
	)
	if err != nil {
		fmt.Println(err)
		panic("Failed to auto-migrate database")
	}

	return db
}
