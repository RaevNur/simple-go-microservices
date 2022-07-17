package helper

import "fmt"

type DbError struct {
	Title       string
	Description string
}

func (e *DbError) Error() string {
	return fmt.Sprintf("%s: %s", e.Title, e.Description)
}
