package model

import (
	v1 "github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
	"gorm.io/gorm"
	"time"
)

type Batch struct {
	gorm.Model
	Name         string `gorm:"unique;not null"`
	Quota        string
	PreDefenceAt time.Time
	DefenceAt    time.Time

	CreatedByID uint  `gorm:"not null"`
	CreatedBy   *User `gorm:"foreignKey:CreatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (b *Batch) TableName() string {
	return "batches"
}

func (b *Batch) convertToMinimalApiFormat() *v1.BatchInfo {
	return &v1.BatchInfo{
		ID:           b.ID,
		Name:         b.Name,
		PreDefenceAt: b.PreDefenceAt,
		DefenceAt:    b.DefenceAt,

		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func (b *Batch) ConvertToApiFormat() *v1.BatchInfo {
	resp := b.convertToMinimalApiFormat()
	if b.CreatedBy != nil {
		resp.CreatedBy = b.CreatedBy.ConvertToMinimalApiFormat()
	}

	return resp
}
