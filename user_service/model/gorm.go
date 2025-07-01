package model

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	CreateAt  *time.Time      `json:"create_at" gorm:"default:null"`
	UpdateAt  *time.Time      `json:"updated_at" gorm:"default:null"`
	DeleteAt  *gorm.DeletedAt `json:"deleted_at" gorm:"index,default:null"`
	IsDeleted bool            `json:"is_deleted"`
}
