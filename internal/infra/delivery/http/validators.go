package http

import (
	"slices"

	"pizza-tracker-go/internal/domain/entity"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valid_pizza_type", sliceValidator(entity.PizzaTypes))
		v.RegisterValidation("valid_pizza_size", sliceValidator(entity.PizzaSizes))
	}
}

func sliceValidator(allowed []string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		return slices.Contains(allowed, fl.Field().String())
	}
}
