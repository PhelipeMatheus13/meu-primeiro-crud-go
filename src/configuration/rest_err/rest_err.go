package rest_err

import "net/http"

// RestErr represents the error object.
// @Summary Error information.
// @Description Structure for describing why the error occurred.
type RestErr struct {
	//Error message.
	Message string `json:"message"`
	//Error description.
	Err string `json:"error"`
	//Error code.
	Code int `json:"code"`
	//Error causes.
	Causes []Causes `json:"causes,omitempty"`
}

// Causes represents the error causes.
// @Summary Error Causes
// @Description Structure representing the causes of an error.
type Causes struct {
	// Field associated with the error cause.
	// @json
	// @jsonTag field
	Field string `json:"field"`

	// Error message describing the cause.
	// @json
	// @jsonTag message
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFounf(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
