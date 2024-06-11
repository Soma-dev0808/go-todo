package repository

import "go_todo/domain/model"

type User interface {
	// Create(u *model.User) error
	Create(user *model.User, grade *model.Grade) error
	Update(u *model.User) error
	Delete(id int) error
	Find(id int) (*model.User, error)
	FindBy(name *string, grade *int) ([]*model.User, error)
	FindAll() ([]*model.User, error)
}