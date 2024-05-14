package repository

import "go_todo/domain/model"

type Todo interface {
	Create(t *model.Todo) error
	Update(t *model.Todo) error
	Delete(id int) error
	Find(id int) (*model.Todo, error)
	FindAll() ([]*model.Todo, error)
}