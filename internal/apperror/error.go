package apperror

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal

}

func NewAppError(c *gin.Context, message string, code int) {
	err := &AppError{
		Message: message,
		Code:    code,
	}
	c.Writer.Header().Set("Content-type", "application/json")
	c.Writer.WriteHeader(code)
	c.Writer.Write(err.Marshal())

}
