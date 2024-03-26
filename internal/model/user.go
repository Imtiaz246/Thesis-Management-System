package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UniversityId string `gorm:"unique;index;primaryKey;not null"`
	Email        string `gorm:"unique;index;not null"`
	Password     string `gorm:"not null"`
	IsAdmin      bool
	Role         Role
	Gender       Gender
	IsVerified   bool
	ChangePass   bool

	Student *Student
	Teacher *Teacher
	Stuff   *Stuff
}

func (u *User) TableName() string {
	return "users"
}

type Role uint8

const (
	RoleStudent Role = iota + 1
	RoleTeacher
	RoleStuff
)

type Gender uint8

const (
	GenderMale Gender = iota + 1
	GenderFemale
)

type Student struct {
	gorm.Model
	Name            string
	Department      string
	CGPA            float64
	Batch           uint8
	Section         string
	Country         string
	Mobile          string
	AlternateMobile string
	UserID          uint `gorm:"unique;not null"`
}

func (s *Student) TableName() string {
	return "students"
}

type Teacher struct {
	gorm.Model
	Name            string
	Department      string
	Designation     string
	Mobile          string
	AlternateMobile string
	UserID          uint `gorm:"unique;not null"`
}

func (t *Teacher) TableName() string {
	return "teachers"
}

type Stuff struct {
	gorm.Model
	Name            string
	Mobile          string
	AlternateMobile string
	UserID          uint `gorm:"unique:not null"`
}

func (s *Stuff) TableName() string {
	return "stuffs"
}
