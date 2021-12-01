package common

import (
	"gorm.io/gorm"
	"time"
)

type SoftControlModel struct {
	Id        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at" sql:"type:timestamptz, default:now()"`
	UpdatedAt time.Time      `json:"updated_at" sql:"type:timestamptz"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
