package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key"`
	Title          string    `gorm:"type:text;not null"`
	Link           string    `gorm:"type:text;uniqueIndex;not null"`
	Source         string    `gorm:"type:varchar(255);not null"`
	Favicon        string    `gorm:"type:text"`
	PublishedAt    time.Time `gorm:"index;not null"`
	ContentSnippet string    `gorm:"type:text"`
	Category       string    `gorm:"type:varchar(100)"`
	ThumbnailURL   string    `gorm:"type:text"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return
}
