package infrastructure

import (
	"go_todo/domain/model"
	"go_todo/domain/repository"

	"gorm.io/gorm"
)

type Recipe struct {
	db *gorm.DB
}

func NewRecipe(db *gorm.DB) repository.Recipe {
	return &Recipe{db: db}
}

func (rd *Recipe) Create(r *model.Recipe) error {
	if err := rd.db.Create(r).Error; err != nil {
		return err
	}
	return nil
}

func (rd *Recipe) Update(r *model.Recipe) error {
	if err := rd.db.Save(r).Error; err != nil {
		return err
	}
	return nil
}

func (rd *Recipe) Delete(id int) error {
	if err := rd.db.Where("id = ?", id).Delete(&model.Recipe{}).Error; err != nil {
		return err
	}
	return nil
}

func (rd *Recipe) Find(id int) (*model.Recipe, error) {
	var recipe *model.Recipe
	err := rd.db.Where("id = ?", id).Take(&recipe).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return recipe, nil
}

func (rd *Recipe) FindAll() ([]*model.Recipe, error) {
	var recipes []*model.Recipe
	err := rd.db.Find(&recipes).Error
	if err != nil {
		return nil, err
	}
	return recipes, nil
}