package models

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Article struct {
	ID             string    `gorm:"type:varchar(26);primary_key"`
	Title          string    `gorm:"type:text;not null;index:idx_article_title,type:gin,expression:title gin_trgm_ops"`
	Link           string    `gorm:"type:text;uniqueIndex;not null"`
	Source         string    `gorm:"type:varchar(255);not null"`
	Favicon        string    `gorm:"type:text"`
	PublishedAt    time.Time `gorm:"index;not null"`
	ContentSnippet string    `gorm:"type:text"`
	Category       string    `gorm:"type:varchar(100);index"`
	ThumbnailURL   string    `gorm:"type:text"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == "" {
		t := a.PublishedAt
		if t.IsZero() {
			t = time.Now()
		}
		entropy := ulid.Monotonic(rand.Reader, 0)
		a.ID = ulid.MustNew(ulid.Timestamp(t), entropy).String()
	}
	return
}

