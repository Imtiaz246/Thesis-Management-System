package model

import (
	v1 "github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UniversityId string `gorm:"unique;index;primaryKey;not null"`
	Email        string `gorm:"unique;index;not null"`
	Password     string `gorm:"not null"`
	IsAdmin      bool
	Role         role   `gorm:"default:1"`
	Gender       gender `gorm:"default:1"`
	IsVerified   bool   `gorm:"default:false"`
	ChangePass   bool   `gorm:"default:false"`

	Student *Student
	Teacher *Teacher
	Stuff   *Stuff
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) ConvertToMinimalApiFormat() *v1.UserResponseData {
	resp := &v1.UserResponseData{
		UniversityId: u.UniversityId,
		Email:        u.Email,
		Role:         u.Role.String(),
		Gender:       u.Gender.String(),
	}
	switch {
	case u.Role == RoleStudent && u.Student != nil:
		resp.Student = u.Student.convertToMinimalApiFormat()
	case u.Role == RoleTeacher && u.Teacher != nil:
		resp.Teacher = u.Teacher.convertToApiFormat()
	case u.Role == RoleStuff && u.Stuff != nil:
		resp.Stuff = u.Stuff.convertToApiFormat()
	default:
	}

	return resp
}

func (u *User) ConvertToApiFormat() *v1.UserResponseData {
	resp := &v1.UserResponseData{
		UniversityId: u.UniversityId,
		Email:        u.Email,
		IsAdmin:      u.IsAdmin,
		IsVerified:   u.IsVerified,
		ChangePass:   u.ChangePass,
		Role:         u.Role.String(),
		Gender:       u.Gender.String(),
	}
	switch {
	case u.Role == RoleStudent && u.Student != nil:
		resp.Student = u.Student.convertToApiFormat()
	case u.Role == RoleTeacher && u.Teacher != nil:
		resp.Teacher = u.Teacher.convertToApiFormat()
	case u.Role == RoleStuff && u.Stuff != nil:
		resp.Stuff = u.Stuff.convertToApiFormat()
	default:
	}

	return resp
}

type role uint8

const (
	RoleStudent role = iota + 1
	RoleTeacher
	RoleStuff
)

func (r role) String() string {
	return []string{"Student", "Teacher", "Stuff"}[r-1]
}

type gender uint8

const (
	GenderMale gender = iota + 1
	GenderFemale
)

func (g gender) String() string {
	return []string{"male", "female"}[g-1]
}

type Student struct {
	gorm.Model
	Name             string
	Department       string
	CGPA             float64
	CompletedCredits uint16
	Batch            uint16
	Section          string
	Country          string
	Mobile           string
	AlterMobile      string
	UserID           uint `gorm:"unique;not null"`
}

func (s *Student) TableName() string {
	return "students"
}

func (s *Student) convertToMinimalApiFormat() *v1.StudentInfo {
	return &v1.StudentInfo{
		FullName:   s.Name,
		Country:    s.Country,
		Department: s.Department,
		Batch:      s.Batch,
	}
}

func (s *Student) convertToApiFormat() *v1.StudentInfo {
	return &v1.StudentInfo{
		FullName:         s.Name,
		Country:          s.Country,
		Department:       s.Department,
		CGPA:             s.CGPA,
		Batch:            s.Batch,
		Section:          s.Section,
		CompletedCredits: s.CompletedCredits,
		Mobile:           s.Mobile,
		AlternateMobile:  s.AlterMobile,
	}
}

type Teacher struct {
	gorm.Model
	Name        string
	Department  string
	Designation string
	Mobile      string
	AlterMobile string
	UserID      uint `gorm:"unique;not null"`
}

func (t *Teacher) TableName() string {
	return "teachers"
}

func (t *Teacher) convertToApiFormat() *v1.TeacherInfo {
	return &v1.TeacherInfo{
		FullName:        t.Name,
		Department:      t.Department,
		Designation:     t.Designation,
		Mobile:          t.Mobile,
		AlternateMobile: t.AlterMobile,
	}
}

type Stuff struct {
	gorm.Model
	Name        string
	Mobile      string
	AlterMobile string
	Department  string
	UserID      uint `gorm:"unique:not null"`
}

func (s *Stuff) TableName() string {
	return "stuffs"
}

func (s *Stuff) convertToApiFormat() *v1.StuffInfo {
	return &v1.StuffInfo{
		FullName:        s.Name,
		Department:      s.Department,
		Mobile:          s.Mobile,
		AlternateMobile: s.AlterMobile,
	}
}
