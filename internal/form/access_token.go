package form

import (
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
)

type DeleteAccessTokenOption struct {
	Token string `json:"token" binding:"Required"`
}

func (f *DeleteAccessTokenOption) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}
