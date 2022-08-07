package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        uint   `gorm:"autoIncrement;primaryKey"`
	Name      string `gorm:"varchar(10);not null;unique"`
	Password  string `gorm:"type:varchar(32);not null"`
	Gender    bool
	Picture   string `gorm:"type:varchar(100)"`
	Email     string `gorm:"type:varchar(40)"`
	RoleId    byte   `gorm:"default:1"`
	Intro     string `gorm:"type:varchar(80)"`
	LastLogin time.Time
}
