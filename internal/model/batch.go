package model

import (
	v1 "github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
	"gorm.io/gorm"
	"time"
)

type Batch struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Quota string

	TeamRegDeadline time.Time // TeamRegDeadline is the cutoff date/time by which team registration must be completed.
	MaxTeamMember   uint8     // MaxTeamMember defines the maximum number of members allowed in a team.
	MaxTeacherPref  uint8     // MaxTeacherPref indicates the maximum number of teacher selections that a team can list as preferences.

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
		ID:    b.ID,
		Name:  b.Name,
		Quota: b.Quota,

		TeamRegDeadline: b.TeamRegDeadline,
		MaxTeamMember:   b.MaxTeamMember,
		MaxTeacherPref:  b.MaxTeacherPref,
		PreDefenceAt:    b.PreDefenceAt,
		DefenceAt:       b.DefenceAt,

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
