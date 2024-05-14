package handler

import (
	"go_todo/domain/model"
	"go_todo/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Recipe interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Find(c *gin.Context)
	FindAll(c *gin.Context)
}

type recipeHandler struct {
	recipeUsecase usecase.Recipe
}

func NewRecipe(u usecase.Recipe) Recipe {
	return &recipeHandler{recipeUsecase: u}
}

type CreateRecipeRequestParam struct {
	Name string `json:"name" binding:"required,max=60"`
}

func (r *recipeHandler) Create(c *gin.Context) {
	var req CreateRecipeRequestParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	if err := r.recipeUsecase.Create(req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

type UpdateRecipeRequestParam struct {
	ID int `uri:"id"`
}

type UpdateRecipeRequestBodyParam struct {
	Name string				`json:"name" binding:"required,max=60"`
	Type model.RecipeType	`json:"type" binding:"required,recipe_type"`
}

// TODO: 柔軟にするNameのみの場合やTypeのみの場合を考慮
func (r *recipeHandler) Update(c *gin.Context) {
	var pathParam UpdateRecipeRequestParam
	var bodyParam UpdateRecipeRequestBodyParam

	if err := c.ShouldBindUri(&pathParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&bodyParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	if err := r.recipeUsecase.Update(pathParam.ID, bodyParam.Name, bodyParam.Type); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result":"success"})
}
type DeleteRecipeRequestParam struct {
	ID int `uri:"id"`
}

func (r *recipeHandler) Delete(c *gin.Context) {
	var req DeleteRecipeRequestParam
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if err := r.recipeUsecase.Delete(req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

type FindRecipeRequestParam struct {
	ID int `uri:"id" binding:"required"`
}

func (r *recipeHandler) Find(c *gin.Context) {
	var req FindRecipeRequestParam
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	recipe, err := r.recipeUsecase.Find(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, recipe)
}

func (r *recipeHandler) FindAll(c *gin.Context) {
	recipes, err := r.recipeUsecase.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, recipes)
} 