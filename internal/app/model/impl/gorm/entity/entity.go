package entity

import (
	"context"
	"time"

	"gin-admin-template/internal/app/config"
	"gin-admin-template/internal/app/icontext"
	"github.com/jinzhu/gorm"
)

// Model base model
type Model struct {
	ID        string     `gorm:"column:id;primary_key;size:36;"`
	CreatedAt time.Time  `gorm:"column:created_at;index;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;index;"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;"`
}

// GetDB ...
func GetDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := icontext.FromTrans(ctx)
	if ok && !icontext.FromNoTrans(ctx) {
		db, ok := trans.(*gorm.DB)
		if ok {
			if icontext.FromTransLock(ctx) {
				if dbType := config.C.Gorm.DBType; dbType == "mysql" ||
					dbType == "postgres" {
					db = db.Set("gorm:query_option", "FOR UPDATE")
				}
			}
			return db
		}
	}
	return defDB
}

// GetDBWithModel ...
func GetDBWithModel(ctx context.Context, defDB *gorm.DB, m interface{}) *gorm.DB {
	return GetDB(ctx, defDB).Model(m)
}
