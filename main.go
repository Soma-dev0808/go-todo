package main

import (
	"fmt"
	"go_todo/handler"
	"go_todo/infrastructure"
	"go_todo/usecase"

	appvalidator "go_todo/handler/validator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	
	db, err := infrastructure.NewDB()
	if err != nil {
		fmt.Printf("failed to start server. db setup failed. err = %s", err.Error())
		return
	}

	r := setupRouter(db)
	if err := appvalidator.SetupValidator(); err != nil {
		fmt.Printf("failed to start server. validator setup failed, err = %s", err.Error())
		return
	}

	r.Run()
}

func setupRouter(d *gorm.DB) *gin.Engine {
	r := gin.Default()

	repository := infrastructure.NewTodo(d)
	todoUsecase := usecase.NewTodo(repository)
	todoHandler := handler.NewTodo(todoUsecase)

	recipeRepository := infrastructure.NewRecipe(d)
	recipeUsecase := usecase.NewRecipe(recipeRepository)
	recipeHandler := handler.NewRecipe(recipeUsecase)

	todo := r.Group("/todo")
	{
		todo.POST("", todoHandler.Create)
		todo.GET("/", todoHandler.FindAll)
		todo.GET("/:id", todoHandler.Find)
		todo.PUT("/:id", todoHandler.Update)
		todo.DELETE("/:id", todoHandler.Delete)
	}

	recipe := r.Group("/recipe")
	{
		recipe.POST("", recipeHandler.Create)
		recipe.GET("/", recipeHandler.FindAll)
		recipe.GET("/:id", recipeHandler.Find)
		recipe.PUT("/:id", recipeHandler.Update)
		recipe.DELETE("/:id", recipeHandler.Delete)
	}

	return r
}