package handler

import (
	// "go_todo/domain/model"
	"go_todo/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Find(c *gin.Context)
	FindAll(c *gin.Context)
}

type UserHandler struct {
	userUsecase usecase.User
}

func NewUserHandler(u usecase.User) User {
	return &UserHandler{userUsecase: u}
}

type CreateUserRequestParam struct {
	Name string `json:"name" binding:"required,max=60"`
	GradeScore int `json:"grade_score" binding:"required,max=60"` // TODO:requiredではない場合の実装を確認
}

func (u *UserHandler) Create(c *gin.Context) {
	var req CreateUserRequestParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if err := u.userUsecase.Create(req.Name, req.GradeScore); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

type UpdateUserRequestParam struct {
	ID int `uri:"id" binding:"required"`
}

// TODO: binding:でmaxを設置。nilの場合に弾かれてしまう。カスタムバリデーション？
type UpdateUserRequestBodyParam struct {
	Name *string `json:"name"`
	GradeScore *int `json:"grade_score"`
}

func (u *UserHandler) Update(c *gin.Context) {
	var reqUri UpdateUserRequestParam
	var reqBody UpdateUserRequestBodyParam

	if err := c.ShouldBindUri(&reqUri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := u.userUsecase.Update(reqUri.ID, reqBody.Name, reqBody.GradeScore); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusNoContent, nil)
}

type DeleteUserRequestParam struct {
	ID int `uri:"id" binding:"required"`
}

func (u *UserHandler) Delete(c *gin.Context) {
	var req FindUserRequestParam
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := u.userUsecase.Delete(req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

type FindUserRequestParam struct {
	ID int `uri:"id" binding:"required"`
}

func (u *UserHandler) Find(c *gin.Context) {
	var req FindUserRequestParam
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := u.userUsecase.Find(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) FindAll(c *gin.Context) {
	users, err := u.userUsecase.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}