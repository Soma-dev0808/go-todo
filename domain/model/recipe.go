package model

import "time"

type Recipe struct {
	ID int `gorm:"primaryKey"`
	Name string
	Type RecipeType
	CreatedAt time.Time `gorm:"<-:false"`
	UpdatedAt time.Time `gorm:"<-:false"`
}

func (Recipe) TableName() string { return "recipes" }

func NewRecipe(name string) *Recipe{
	return &Recipe{
		Name: name, 
		Type: Normal,
	}
}

func NewUpdateRecipe(id int, name string, recipeType RecipeType) *Recipe {
	return &Recipe{
		ID: id,
		Name: name,
		Type: recipeType,
	}
}

type RecipeType string

const (
	Normal	= RecipeType("normal")
	Family 	= RecipeType("family")
	Hurry 	= RecipeType("hurry")
	Easy 	= RecipeType("easy")
)

var RecipeTypeMap = map[RecipeType]bool {
	Family: true,
	Hurry:  true,
	Easy:   true,
}