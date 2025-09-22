package errcode

import (
	"net/http"

	"github.com/skyrocket-qy/erx"
)

type Err string

func (c Err) Str() string {
	return string(c)
}

const (
	ErrBadRequest               Err = "400.0000"
	ErrEmptyRequest             Err = "400.0001"
	ErrParsePayload             Err = "400.0002"
	ErrValidateInput            Err = "400.0003"
	ErrAlreadyResetOTP          Err = "400.0004"
	ErrUnsupportedOAuthProvider Err = "400.0005"

	ErrUnauthorized               Err = "401.0000"
	ErrNewPasswordRequired        Err = "401.0001"
	ErrMissingAuthorizationHeader Err = "401.0002"

	ErrNotFound Err = "404.0000"

	ErrDuplicate     Err = "409.0000"
	ErrUnknown       Err = "500.0000"
	ErrDBUnavailable Err = "500.0001"
	ErrLokiError     Err = "500.0002"
	ErrNoWrapErx     Err = "500.0003"

	ErrNotImplemented Err = "501.0000"
)

var ErrToHTTP = map[erx.Code]int{
	ErrBadRequest:      http.StatusBadRequest,
	ErrEmptyRequest:    http.StatusBadRequest,
	ErrParsePayload:    http.StatusBadRequest,
	ErrValidateInput:   http.StatusBadRequest,
	ErrAlreadyResetOTP: http.StatusBadRequest,

	ErrUnauthorized:               http.StatusUnauthorized,
	ErrNewPasswordRequired:        http.StatusUnauthorized,
	ErrMissingAuthorizationHeader: http.StatusUnauthorized,

	ErrNotFound: http.StatusNotFound,

	ErrUnknown:       http.StatusInternalServerError,
	ErrDBUnavailable: http.StatusInternalServerError,

	ErrNotImplemented: http.StatusNotImplemented,
}
