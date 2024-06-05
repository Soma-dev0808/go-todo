package model

import (
	"time"
)

type User struct {
	ID int `gorm:"primaryKey"`
	Name string
	CreatedAt time.Time `gorm:"<-:false"`
	UpdatedAt time.Time `gorm:"<-:false"`
	// TODO：constraint:OnDelete:CASCADE;が動かない
	// TODO: ポインタである必要？ => ポインタ型ではない場合、nullで表示されない。ゼロを値にもつ構造体になる
	Grade	*Grade `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"` // has oneの関係
}

// テーブル名？
func (User) TableName() string { return "users" }

func NewUser(name string, gradeScore int) *User {
	return &User{
		Name: name,
		Grade: &Grade{
			Score: gradeScore,
		},
	}
}