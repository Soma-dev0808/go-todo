package model

import "time"

type User struct {
	ID int `gorm:"primaryKey"`
	Name string
	CreatedAt time.Time `gorm:"<-:false"`
	UpdatedAt time.Time `gorm:"<-:false"`
	Grade	Grade `gorm:"foreignKey:UserId;references:ID"`
}


func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}