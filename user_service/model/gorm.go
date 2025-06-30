package model

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	CreateAt  time.Time      `json:"create_at"`
	UpdateAt  time.Time      `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	IsDeleted bool           `json:"is_deleted"`
}
