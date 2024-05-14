package usecase

import (
	"go_todo/domain/model"
	"go_todo/domain/repository"
)

type Recipe interface {
	Create(name string) error
	Update(id int, name string, recipeType model.RecipeType) error
	Delete(id int) error
	Find(id int) (*model.Recipe, error)
	FindAll() ([]*model.Recipe, error)
}

type recipe struct {
	recipeRepository repository.Recipe
}

func NewRecipe(r repository.Recipe) Recipe {
	return &recipe{recipeRepository: r}
}

func (r *recipe) Create(name string) error {
	newRecipe := model.NewRecipe(name)
	if err := r.recipeRepository.Create(newRecipe); err != nil {
		return err
	}
	return nil
}

func (r *recipe) Update(id int, name string, recipeType model.RecipeType) error {
	newRecipe := model.NewUpdateRecipe(id, name, recipeType)
	if err := r.recipeRepository.Update(newRecipe); err != nil {
		return err
	}
	return nil
}

func (r *recipe) Delete(id int) error {
	if err := r.recipeRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (r *recipe) Find(id int) (*model.Recipe, error) {
	recipe, err := r.recipeRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func (r *recipe) FindAll() ([]*model.Recipe, error) {
	recipes, err := r.recipeRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return recipes, nil
}