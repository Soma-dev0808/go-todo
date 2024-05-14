package model

type Grade struct {
	ID int `gorm:"primaryKey"`
	UserID int
	Score int
}