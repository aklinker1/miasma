package database

import (
	"github.com/aklinker1/miasma/internal/shared/log"
	"gorm.io/gorm"
)

func TX(err *error) (tx *gorm.DB, onDefer func()) {
	db := open()
	tx = db.Begin()
	onDefer = func() {
		if panicErr := recover(); panicErr != nil {
			log.E("Paniced, rolling back: %v", panicErr)
			tx.Rollback()
		}
		if errVal := *err; errVal != nil {
			log.E("Request failed, rolling back (%v)", errVal)
			tx.Rollback()
			return
		}
		tx.Commit()
	}
	return
}
