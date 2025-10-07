package api_err

import (
	"log"
	"time"
)

type ApiError struct {
	Code       string
	Message    string
	HTTPStatus int
	Err        error
}

type ApiErrorResponse struct {
	Code          string `json:"code"`
	Message       string `json:"message"`
	CorrelationId string `json:"correlationId"`
	ServiceID     string `json:"serviceId"`
	Timestamp     string `json:"timestamp"`
}

func (e *ApiError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func (e *ApiError) WithMessage(msg string) *ApiError {
	cp := *e
	cp.Message = msg
	return &cp
}

func (e *ApiError) Clone() *ApiError {
	cp := *e
	return &cp
}

func (e *ApiError) ToJSON(serviceID string, correlationID string) *ApiErrorResponse {
	return &ApiErrorResponse{
		Code:          e.Code,
		Message:       e.Message,
		CorrelationId: correlationID,
		ServiceID:     serviceID,
		Timestamp:     time.Now().Format(time.RFC3339),
	}
}

func (e *ApiError) Log() {
	log.Printf("[ERROR] %d %s \"%s\": %v", e.HTTPStatus, e.Code, e.Message, e.Error())
}

func (e *ApiError) Unwrap() error {
	return e.Err
}

func Unknown(err error) *ApiError {
	apierr := ErrInternalServer.Clone()
	apierr.Err = err
	return apierr

}

func Wrap(err *ApiError, internal error) *ApiError {
	if internal == nil {
		return err
	}

	return &ApiError{
		Code:       err.Code,
		Message:    err.Message,
		HTTPStatus: err.HTTPStatus,
		Err:        internal,
	}
}
