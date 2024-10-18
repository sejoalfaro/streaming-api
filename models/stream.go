package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Stream struct {
	ID        string         `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    string         `gorm:"type:uuid;not null" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID"`
	StreamKey string         `gorm:"uniqueIndex;not null" json:"stream_key"`
	RTMPURL   string         `gorm:"not null" json:"rtmp_url"`
	HLSURL    string         `gorm:"not null" json:"hls_url"`
	Status    string         `gorm:"default:active" json:"status"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (s *Stream) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return
}
