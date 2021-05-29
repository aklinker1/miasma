package database

import (
	"database/sql"

	"gorm.io/gorm"
)

func ReadOnly() (tx *gorm.DB, onDefer func()) {
	db := open()
	tx = db.Begin(&sql.TxOptions{
		ReadOnly: true,
	})
	onDefer = func() {
		tx.Rollback()
	}
	return
}
