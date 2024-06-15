package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key:unique" json:"id"`
	FirstName    string    `gorm:"type:varchar(255);size:20;not null" json:"fname"`
	LastName     string    `gorm:"type:varchar(255);size:20;not null" json:"lname"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	Password     string    `gorm:"size:500;not null" json:"password,omitempty"`
	ProfilePhoto string    `gorm:"type:varchar(255);not null" json:"profile_photo"`
	Verification int64     `gorm:"size:20;default:0;not null" json:"verification,omitempty"`
	Verified     bool      `gorm:"not null" json:"verified,omitempty"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Login struct {
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
