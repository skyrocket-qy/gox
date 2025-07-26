package errcode

import (
	"net/http"

	"github.com/skyrocket-qy/gox/erx"
)

type err string

func (c err) Str() string {
	return string(c)
}

const (
	ErrBadRequest               err = "400.0000"
	ErrEmptyRequest             err = "400.0001"
	ErrParsePayload             err = "400.0002"
	ErrValidateInput            err = "400.0003"
	ErrAlreadyResetOTP          err = "400.0004"
	ErrUnsupportedOAuthProvider err = "400.0005"

	ErrUnauthorized               err = "401.0000"
	ErrNewPasswordRequired        err = "401.0001"
	ErrMissingAuthorizationHeader err = "401.0002"

	ErrNotFound err = "404.0000"

	ErrDuplicate     err = "409.0000"
	ErrUnknown       err = "500.0000"
	ErrDBUnavailable err = "500.0001"
	ErrLokiError     err = "500.0002"
	ErrNoWrapErx     err = "500.0003"

	ErrNotImplemented err = "501.0000"
)

var errToHTTP = map[erx.Code]int{
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
