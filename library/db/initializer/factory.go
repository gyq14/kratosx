package initializer

import (
	"github.com/gyq14/kratosx/library/db/initializer/impl"
	"github.com/gyq14/kratosx/library/db/model"
	"gorm.io/gorm"
)

// New 工厂方法
func New(dbType model.DatabaseType, db *gorm.DB, path string, force bool) model.Initializer {
	if err := db.AutoMigrate(&model.GormInit{}); err != nil {
		panic("failed to migrate GormInit table: " + err.Error())
	}

	switch dbType {
	case model.Postgres:
		return impl.NewPgSQLInitializer(db, path, force)
	default: // 默认MySQL
		return impl.NewMySQLInitializer(db, path, force)
	}
}
