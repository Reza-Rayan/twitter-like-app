package models

type Role struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}
