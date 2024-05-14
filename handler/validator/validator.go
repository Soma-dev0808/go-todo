package validator

import (
	"go_todo/domain/model"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// handlerで使用されるカスタムのバリデーション
func SetupValidator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("task_status", ValidateTaskStatus); err != nil {
			return err
		}
		if err := v.RegisterValidation("recipe_type", ValidateRecipeType); err != nil {
			return err
		}
	}
	return nil
}

func ValidateTaskStatus(fl validator.FieldLevel) bool{
	return model.TaskStatusMap[model.TaskStatus(fl.Field().String())]
}

func ValidateRecipeType(fl validator.FieldLevel) bool{
	return model.RecipeTypeMap[model.RecipeType(fl.Field().String())]
}