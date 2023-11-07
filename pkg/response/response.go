package response

import (
	"github.com/gin-gonic/gin"
)

type MsgResponse struct {
	Message string `json:"msg"`
}

func ErrResponse(c *gin.Context, status int, msg string) error {
	c.JSON(status, &MsgResponse{
		Message: msg,
	})
	return nil
}

func SuccessResponse(c *gin.Context, status int, data interface{}) error {
	c.JSON(status, data)
	return nil
}
