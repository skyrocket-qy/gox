package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/skyrocket-qy/erx"
	"github.com/skyrocket-qy/gox/errcode"
)

var v *validator.Validate

func New() {

	v = validator.New(validator.WithRequiredStructEnabled())
}

func Struct(st any) error {
	if err := v.Struct(st); err != nil {
		return erx.W(err).SetCode(errcode.ErrValidateInput)
	}

	return nil
}
