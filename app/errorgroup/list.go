package errorgroup

import "net/http"

var (
	InternalServerError = Error{
		Code:    500,
		Message: "INTERNAL_SERVER_ERROR",
	}
	UNAUTHORIZED = Error{
		Code:    http.StatusUnauthorized,
		Message: "UNAUTHORIZED",
	}
	TOKEN_EXPIRED = Error{
		Code:    http.StatusUnauthorized,
		Message: "TOKEN_EXPIRED",
	}
	TOKEN_INVALID = Error{
		Code:    http.StatusBadRequest,
		Message: "TOKEN_INVALID",
	}
	HEADER_PAYLOAD_NOT_ALLOWED = Error{
		Code:    http.StatusBadRequest,
		Message: "HEADER_PAYLOAD_NOT_ALLOWED",
	}
	USER_NOT_FOUND = Error{
		Code:    http.StatusNotFound,
		Message: "USER_NOT_FOUND",
	}
	BAD_REQUEST = Error{
		Code:    http.StatusBadRequest,
		Message: "BAD_REQUEST",
	}
)
