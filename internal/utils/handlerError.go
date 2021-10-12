package utils

import "fmt"

type ErrorResponse struct {
	Field   string
	Message string
}

func HandlerError(kind string, field string) string {
	switch kind {
	case "required":
		return fmt.Sprintf("Trường %s là bắt buộc", field)
	default:
		return fmt.Sprintf("Trường %s không đúng định dạng", field)
	}
}
