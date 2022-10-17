package models

import (
	"time"
)

type Members struct {
	Id         int       `gorm:"column:id;primary_key;"`
	FullName   string    `gorm:"column:fullName;not null"`
	Username   string    `gorm:"column:username;NOT NULL"`
	Password   string    `gorm:"column:password;"`
	DateOfBird time.Time `gorm:"column:dateOfBird;type:timestamp;"`
	IsVip      int       `gorm:"column:isVip;not null"`
	UpdatedAt  string    `gorm:"column:updatedAt;type:timestamp"`
	CreatedAt  time.Time `gorm:"column:createdAt;type:timestamp;default:current_timestamp"`
}
