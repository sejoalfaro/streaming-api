package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            string         `gorm:"type:uuid;primaryKey" json:"id"`
	Email         string         `gorm:"uniqueIndex;not null" json:"email" binding:"required,email"`
	Name          string         `json:"name" binding:"required"`
	OAuthProvider string         `json:"oauth_provider" binding:"required"`
	OAuthID       string         `gorm:"uniqueIndex" json:"oauth_id"`
	Streams       []Stream       `gorm:"foreignKey:UserID"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
