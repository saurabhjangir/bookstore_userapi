package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status int `json:"status"`
	Error string `json:"error"`
}

func NewRestErrBadRequest(message string) *RestErr{
	return &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "resource_already_exist",
	}
}

func NewRestErrInteralServer(message string) *RestErr{
	return &RestErr{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "internal_server_error",
	}
}

func NewRestErrResourceNotFound(message string) *RestErr{
	return &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "resource_not_found",
	}
}


