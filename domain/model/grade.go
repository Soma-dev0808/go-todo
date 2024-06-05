package model

import "time"

type Grade struct {
	ID int `gorm:"primaryKey"`
	UserID int
	Score int
	CreatedAt time.Time `gorm:"<-:false"`
	UpdatedAt time.Time `gorm:"<-:false"`
}

func (Grade) TableName() string { return "grades" }

func NewGrade(userId int, score int) *Grade{
	return &Grade{
		UserID: userId,
		Score: score,
	}
}