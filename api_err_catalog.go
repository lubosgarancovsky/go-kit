package go_kit

import "net/http"

var (
	ErrBadRequest = &ApiError{
		Code:       "BAD REQUEST",
		Message:    "Bad Request",
		HTTPStatus: http.StatusBadRequest, // 400
	}

	ErrUnauthorized = &ApiError{
		Code:       "UNAUTHORIZED",
		Message:    "Authentication required.",
		HTTPStatus: http.StatusUnauthorized, // 401
	}

	ErrForbidden = &ApiError{
		Code:       "FORBIDDEN",
		Message:    "Access denied",
		HTTPStatus: http.StatusForbidden, // 403
	}

	ErrNotFound = &ApiError{
		Code:       "NOT_FOUND",
		Message:    "Resource not found.",
		HTTPStatus: http.StatusNotFound, //404
	}

	ErrMethodNotAllowed = &ApiError{
		Code:       "METHOD_NOT_ALLOWED",
		Message:    "Http method not allowed.",
		HTTPStatus: http.StatusMethodNotAllowed, //404
	}

	ErrInternalServer = &ApiError{
		Code:       "INTERNAL_SERVER_ERROR",
		Message:    "Internal server error",
		HTTPStatus: http.StatusInternalServerError, //500
	}

	ErrNotImplemented = &ApiError{
		Code:       "NOT_IMPLEMENTED",
		Message:    "Not implemented",
		HTTPStatus: http.StatusNotImplemented, //501
	}

	ErrParameterMissing = &ApiError{
		Code:       "PARAMETER_MISSING",
		Message:    "Parameter missing",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrInvalidUUID = &ApiError{
		Code:       "INVALID_UUID",
		Message:    "Invalid UUID",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrTooManyRequests = &ApiError{
		Code:       "TOO_MANY_REQUESTS",
		Message:    "Too many requests",
		HTTPStatus: http.StatusTooManyRequests, //429
	}
)
