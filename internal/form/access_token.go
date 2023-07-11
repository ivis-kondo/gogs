package form

import (
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
)

type CreateAccessTokenOption struct {
	Name          string `json:"name" binding:"Required"`
	ExpireMinutes int64  `json:"expire_minutes"`
}

func (f *CreateAccessTokenOption) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}
