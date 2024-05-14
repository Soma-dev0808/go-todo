package repository

import "go_todo/domain/model"

type Recipe interface {
	Create(r *model.Recipe) error
	Update(r *model.Recipe) error
	Delete(id int) error
	Find(id int) (*model.Recipe, error)
	FindAll() ([]*model.Recipe, error)
}