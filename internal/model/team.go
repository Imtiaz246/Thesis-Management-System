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

func (_ *TeamMember) TableName() string {
	return "team_members"
}

type TeamTeacherRole uint8

const (
	TeamTeacherRoleNotSelected TeamTeacherRole = iota
	TeamTeacherRoleSupervisor
	TeamTeacherRoleCoSupervisor
)

type SelectionParams struct {
	SelectedRole      TeamTeacherRole `json:"selectedRole"`
	TeamRankByTeacher uint8           `json:"teamRankByTeacher"`
	TeacherRankByTeam uint8           `json:"teacherRankByTeam"`
}

func (sp *SelectionParams) GetSelectedRole() TeamTeacherRole {
	return sp.SelectedRole
}

type TeamTeacher struct {
	gorm.Model

	TeamID    uint     `gorm:"not null;uniqueIndex:idx_team_teacher"`
	TeacherID uint     `gorm:"not null;uniqueIndex:idx_team_teacher"`
	Team      *Team    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Teacher   *Teacher `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	SelectionParams SelectionParams `gorm:"type:jsonb;" json:"selectionParams"`
}

func (_ *TeamTeacher) TableName() string {
	return "team_teachers"
}

type TeamInvitation struct {
	gorm.Model

	TeamID    uint     `gorm:"not null;uniqueIndex:idx_team_student"`
	StudentID uint     `gorm:"not null;uniqueIndex;idx_team_student"`
	Team      *Team    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Student   *Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (_ *TeamInvitation) TableName() string {
	return "team_invitations"
}
