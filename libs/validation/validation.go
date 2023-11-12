package validation

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go-web/utils"
	"net/http"
	"net/url"
)

type Body map[string][]validation.Rule
type ValidationError struct {
	Path    interface{}
	Message string
}
type ValidationErrors []ValidationError

// WithRequestInputValidation
//
// @shouldInterceptRequest: If true, the request will be intercepted and validation errors will be returned immediately. Otherwise, the validation errors will be passed down the middleware into original hanlder func /
type WithRequestInputValidation struct {
	ShouldInterceptRequest bool
}

type WithRequestInputValidationContext struct {
	ValidationErrors ValidationErrors
	FormData         url.Values
}

func (middleware *WithRequestInputValidation) Validate(validationsMap map[string][]validation.Rule) func(next http.Handler) http.Handler {
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
			var vErrs ValidationErrors
			for key, rules := range validationsMap {
				value := form.Get(key)
				if vErr := validation.Validate(value, rules...); vErr != nil {
					ve := ValidationError{Path: key, Message: vErr.Error()}
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
