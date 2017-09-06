package models

import "time"

type baseModel struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	CreatorID   uint
	UpdaterID   uint
	DeleterID   *uint
	LockVersion uint
}
