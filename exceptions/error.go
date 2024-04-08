package exceptions

import (
	"fmt"
)

type Exception struct {
	message    string
	statusCode int
}

func (e *Exception) Error() string {
	return fmt.Sprintf("%s (Status Code: %d)", e.message, e.statusCode)
}

func (e *Exception) GetCode() int                  { return e.statusCode }
func (e *Exception) GetMessage() map[string]string { return map[string]string{"error": e.message} }

func NewError(message string, statusCode int) *Exception {
	if message == "" {
		message = "An error occurred"
	}
	if statusCode == 0 {
		statusCode = 500
	}
	return &Exception{
		message:    message,
		statusCode: statusCode,
	}
}
