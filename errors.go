package commons

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//BadRequestError
type BadRequestError struct {
	Message string
}

//NotFoundError
type NotFoundError struct {
	Message string
}

//InternalError
type InternalError struct {
	Message string
}

func (e *BadRequestError) Error() string {
	return fmt.Sprintf("BadRequest: %v", e.Message)
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("NotFound: %v", e.Message)
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("InternalServerError: %v", e.Message)
}

//ErrorWrapper main struct for custom error return
type ErrorWrapper struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//DecodeError returns a specific type of error according to type
func DecodeError(err error) (status int, errBody ErrorWrapper) {
	switch err.(type) {
	case *BadRequestError:
		return http.StatusBadRequest, ErrorWrapper{Code: http.StatusBadRequest, Message: err.Error()}
	case *NotFoundError:
		return http.StatusNotFound, ErrorWrapper{Code: http.StatusNotFound, Message: err.Error()}
	default:
		return http.StatusInternalServerError, ErrorWrapper{Code: http.StatusInternalServerError, Message: err.Error()}
	}
}

// HandleError handle the custom errors to be returned to the user
func HandleError(w http.ResponseWriter, e error) {
	status, err := DecodeError(e)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}
