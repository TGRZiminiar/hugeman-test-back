package request

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	contextWrapperService interface {
		Bind(data any) error
	}

	contextWrapper struct {
		Context   *gin.Context
		validator *validator.Validate
	}
)

func NewContextWrapper(ctx *gin.Context) contextWrapperService {
	return &contextWrapper{
		Context:   ctx,
		validator: validator.New(),
	}
}

func (c *contextWrapper) Bind(data any) error {
	if err := c.Context.Bind(data); err != nil {
		log.Printf("Error: Bind data failed: %s", err.Error())
		return fmt.Errorf("error: bind data failed: %s", err.Error())
	}

	if err := c.validator.Struct(data); err != nil {
		// log.Printf("Error: Validate data failed: %s", err.Error())
		return fmt.Errorf("error: validate data failed: %s", err.Error())
	}

	return nil
}
