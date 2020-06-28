package objects

import "fmt"

// NotFoundError that will be mapped to 404
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf(e.Message)
}
