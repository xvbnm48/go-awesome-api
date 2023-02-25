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

type Role struct {
	ID   int    `gorm:"column:id; primary_key; not null" json:"id"`
	Role string `gorm:"column:role" json:"role"`
	BaseModel
}

type User struct {
	ID       int    `gorm:"column:id; primary_key; not null" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password;->false" json:"-"`
	Status   int    `gorm:"column:status" json:"status"`
	RoleID   int    `gorm:"column:role_id; not null" json:"role_id"`
	Role     Role   `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	BaseModel
}
