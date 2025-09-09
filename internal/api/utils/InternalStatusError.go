package utils

import (
	"fmt"
	"net/http"
)

type InternalStatusError struct {
	Message string
}

func (e *InternalStatusError) Error() string {
	return e.Message
}

func CreateInternalStatusErrorWithMessage(message string) *InternalStatusError {
	return &InternalStatusError{
		Message: message,
	}
}
func CreateInternalStatusError() *InternalStatusError {
	return &InternalStatusError{
		Message: "",
	}
}
func (e *InternalStatusError) Write(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	if e.Message != "" {
		fmt.Fprintf(w, "{\"error\": \"%s\"}", e.Message)
	} else {
		fmt.Fprintf(w, "{\"error\": \"Internal Server Error\"}")
	}
	e.Message = "Internal Server Error"
}
