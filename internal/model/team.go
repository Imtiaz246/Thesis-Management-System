package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model

	Name    string `gorm:"not null"`
	Subject string `gorm:"not null"`

	BatchID uint   `gorm:"not null"`
	Batch   *Batch `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Members  []TeamMember
	Teachers []TeamTeacher
}

func (t *Team) TableName() string {
	return "teams"
}

type TeamMember struct {
	gorm.Model

	TeamID    uint     `gorm:"not null;uniqueIndex:idx_team_student"`
	StudentID uint     `gorm:"not null;uniqueIndex:idx_team_student"`
	Team      *Team    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Student   *Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (tm *TeamMember) TableName() string {
	return "team_members"
}

type TeamTeacherRole uint8

const (
	TeamTeacherRoleNotSelected TeamTeacherRole = iota
	TeamTeacherRoleSupervisor
	TeamTeacherRoleCoSupervisor
)

type TeamTeacher struct {
	gorm.Model

	TeamID    uint     `gorm:"not null;uniqueIndex:idx_team_teacher"`
	TeacherID uint     `gorm:"not null;uniqueIndex:idx_team_teacher"`
	Team      *Team    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Teacher   *Teacher `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Role TeamTeacherRole `gorm:"not null;default:0"`
	Rank uint8           `gorm:"not null"`
}

func (tt *TeamTeacher) TableName() string {
	return "team_teachers"
}
