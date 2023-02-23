package dao

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	CreateAt  time.Time      `gorm:"->:false;column:create_at" json:"-"`
	UpdatetAt time.Time      `gorn:"->:false;column:update_at" json:"-"`
	DeleteAt  gorm.DeletedAt `gorm:"->:false;column:delete_at" json:"-"`
}
