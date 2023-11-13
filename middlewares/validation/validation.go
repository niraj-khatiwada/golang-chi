package middlewares

import (
	"context"
	validator "github.com/go-ozzo/ozzo-validation/v4"
	"go-web/libs/validation"
	"go-web/utils"
	"net/http"
	"net/url"
)

// WithRequestInputValidation
/*
	- @param {bool} shouldInterceptRequest: If true, the request will be intercepted and validation errors will be returned immediately.: If true, the request will be intercepted and validation errors will be returned immediately. Otherwise, the validation errors will be passed down the middleware into original handler func
*/
type WithRequestInputValidation struct {
	ShouldInterceptRequest bool
}

type WithRequestInputValidationContext struct {
	ValidationErrors validation.ValidationErrors
	FormData         url.Values
}

func (middleware *WithRequestInputValidation) Validate(validationsMap map[string][]validator.Rule) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if validationsMap == nil {
				next.ServeHTTP(w, r)
				return
			}
			if err := r.ParseForm(); err != nil {
				utils.CatchRuntimeErrors(err)
				http.Error(w, "err", 500)
				return
			}
			form := r.PostForm
			c := WithRequestInputValidationContext{}
			c.FormData = form
			var vErrs validation.ValidationErrors
			for key, rules := range validationsMap {
				value := form.Get(key)
				if vErr := validator.Validate(value, rules...); vErr != nil {
					ve := validation.ValidationError{Path: key, Message: vErr.Error()}
					vErrs = append(vErrs, ve)
				}
			}
			if len(vErrs) > 0 {
				if middleware.ShouldInterceptRequest {
					http.Error(w, "validation err", 500)
					return
				} else {
					c.ValidationErrors = vErrs
				}
			}
			ctx := context.WithValue(r.Context(), "WithRequestInputValidation", c)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
