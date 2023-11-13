package middlewares

import (
	"context"
	"fmt"
	"go-web/libs/flash"
	redisLib "go-web/libs/redis"
	"go-web/utils"
	"net/http"
	"time"
)

type CSRF struct {
}

type WithCSRFInjectionContext struct {
	CSRFToken string
}

type WithCSRFVerificationContext struct {
	CSRFToken string
}

const (
	WithCSRFInjection    = "WithCSRFInjection"
	WithCSRFVerification = "WithCSRFVerification"
)

func (_ *CSRF) WithCSRFInjection(redis *redisLib.Redis) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			UUID := utils.AlphanumericUUID{}

			uuid, err := UUID.Generate()
			if err != nil {
				utils.CatchRuntimeErrors(err)
				http.Error(w, "error", 500)
				return
			} else {
				if redisErr := redis.Set(fmt.Sprintf("csrf-token-%s", uuid), uuid, 15*time.Minute); redisErr != nil {
					utils.CatchRuntimeErrors(redisErr)
					http.Error(w, "error", 500)
					return
				}
			}
			ctxVal := WithCSRFInjectionContext{CSRFToken: uuid}
			ctx := context.WithValue(r.Context(), WithCSRFInjection, ctxVal)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (_ *CSRF) WithCSRFVerification(redis *redisLib.Redis) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := r.ParseForm(); err != nil {
				utils.CatchRuntimeErrors(err)
				http.Error(w, "err", 500)
				return
			}
			form := r.PostForm
			csrfTokenStr := form.Get("csrf")
			if len(csrfTokenStr) == 0 {
				f := flash.Flash{}
				f.SetFlashMessage(w, flash.FlashError, "Invalid form")
				http.Redirect(w, r, "/contact", 302)
				return
			}
			var csrfToken string
			csrfTokenKey := fmt.Sprintf("csrf-token-%s", csrfTokenStr)
			redisErr := redis.Get(csrfTokenKey, &csrfToken)
			if redisErr != nil {
				f := flash.Flash{}
				f.SetFlashMessage(w, flash.FlashError, "Form submission has expired.")
				http.Redirect(w, r, "/contact", 302)
				return
			}
			ctxVal := WithCSRFVerificationContext{CSRFToken: csrfToken}
			ctx := context.WithValue(r.Context(), WithCSRFVerification, ctxVal)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
