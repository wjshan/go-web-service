package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type JsonDump interface {
	ToJson() string
}
type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	UserName  string
	Password  sql.NullString
	CreatedAt time.Time `gorm:"autoUpdateTime"`
	UpdateAt  time.Time `gorm:"autoCreateTime"`
}

func (u *User) ToJson() (s string, err error) {
	return "", nil
}
