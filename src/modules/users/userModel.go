package users

import "time"

type User struct {
	ID             uint       `gorm:"primaryKey"`
	FullName       string     `gorm:"not null"`
	Email          string     `gorm:"unique;not null"`
	Password       string     `gorm:"not null"`
	Role           string     `gorm:"not null;default:'user'"`
	IsActive       bool       `gorm:"not null;default:false"`
	CodeExpiry     *time.Time `gorm:"default:null"`
	CodeActive     *string    `gorm:"default:null"`
	LastSendCodeAt *time.Time `gorm:"default:null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
