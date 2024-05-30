package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string          `gorm:"unique"`
	Password string          `json:"-"`
	Profile  ProfileResponse `json:"profile"`
	Note     []NoteResponse  `json:"note"`
}

type UserRequest struct {
	Username string `gorm:"required"`
	Password string `gorm:"required"`
}

type Profile struct {
	gorm.Model
	FirstName string
	LastName  string
	UserId    uint
	User      User `json:"user" gorm:"foreignKey:UserId"`
}

type ProfileResponse struct {
	FirstName string
	LastName  string
	UserId    uint `json:"-"`
}

type Note struct {
	gorm.Model
	Title   string `json:"title" gorm:"required"`
	Body    string `json:"body" gorm:"required"`
	DueDate string `json:"due_date"`
	UserId  uint   `json:"user_id"`
	User    User   `json:"user" gorm:"foreignKey:UserId"`
	Tag     []Tag  `json:"tag" gorm:"many2many:note_tag"`
}

type NoteResponse struct {
	Title   string `json:"title"`
	Body    string `json:"body"`
	DueDate string `json:"due_date"`
	UserId  uint   `json:"-"`
}

type Tag struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}

func (NoteResponse) TableName() string {
	return "notes"
}
