package model

import (
	v1 "github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
	"gorm.io/gorm"
	"time"
)

type Batch struct {
	gorm.Model
	Name    string `gorm:"unique;not null"`
	Quota   string
	MinCGPA float64
	MinCH   uint16 `gorm:"not null"`

	TeamRegDeadline time.Time `gorm:"not null"`
	MaxTeamMember   uint8     `gorm:"not null"`
	MaxTeacherPref  uint8     `gorm:"not null"`
	PreDefenceAt    time.Time `gorm:"not null"`
	DefenceAt       time.Time `gorm:"not null"`
	Closed          bool      `gorm:"not null;index;default:false"`

	CreatedByID uint  `gorm:"not null"`
	CreatedBy   *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (_ *Batch) TableName() string {
	return "batches"
}

func (b *Batch) convertToMinimalApiFormat() *v1.BatchInfo {
	return &v1.BatchInfo{
		ID:    b.ID,
		Name:  b.Name,
		Quota: b.Quota,

		MinCGPARequired: b.MinCGPA,
		MinCHRequired:   b.MinCH,
		TeamRegDeadline: b.TeamRegDeadline,
		MaxTeamMember:   b.MaxTeamMember,
		MaxTeacherPref:  b.MaxTeacherPref,
		PreDefenceAt:    b.PreDefenceAt,
		DefenceAt:       b.DefenceAt,
		Closed:          b.Closed,

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

func (b *Batch) VerifyBeforeUpsert() error {
	if b.TeamRegDeadline.After(b.PreDefenceAt) {
		return v1.ErrInvalidTeamRegDeadline
	}
	if b.PreDefenceAt.After(b.DefenceAt) {
		return v1.ErrInvalidPreDefenceDate
	}
	return nil
}

type BatchStage uint8

const (
	StageTeamRegistration BatchStage = iota + 1
	StagePreDefence
	StageDefence
	StageResult
)

func (b *Batch) GetCurrentStage() BatchStage {
	now := time.Now()
	switch {
	case now.Before(b.TeamRegDeadline):
		return StageTeamRegistration
	case now.Before(b.PreDefenceAt):
		return StagePreDefence
	case now.Before(b.DefenceAt):
		return StageDefence
	default:
		return StageResult
	}
}

type BatchRegistration struct {
	gorm.Model

	BatchID   uint     `gorm:"not null;uniqueIndex:idx_batch_student"`
	StudentID uint     `gorm:"not null;uniqueIndex:idx_batch_student"`
	Batch     *Batch   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Student   *Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (_ *BatchRegistration) TableName() string {
	return "batch_registrations"
}
