package csrf

import (
	"context"
	"fmt"
	redisLib "go-web/libs/redis"
	"go-web/utils"
	"net/http"
	"net/url"
	"time"
)

type CSRF struct {
	targetDOMName string
}

const (
	CSRFTokenContextKey = "CSRFToken"
)

func (csrf *CSRF) WithCSRFInjection(redis *redisLib.Redis) func(next http.Handler) http.Handler {
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
			ctx := context.WithValue(r.Context(), CSRFTokenContextKey, uuid)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (csrf *CSRF) WithCSRFVerification(redis *redisLib.Redis) func(next http.Handler) http.Handler {
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
				// TODO: Integrate Session and redirect back to /contact
				http.Redirect(w, r, "/contact?error=invalid form ", 410)
				return
			}
			var csrfToken string
			csrfTokenKey := fmt.Sprintf("csrf-token-%s", csrfTokenStr)
			redisErr := redis.Get(csrfTokenKey, &csrfToken)
			if redisErr != nil {
				// TODO: Integrate Session and redirect back to /contact
				w.Header().Set("Content-Type", "")
				http.Redirect(w, r, fmt.Sprintf("/contact?error=%s", url.QueryEscape("form submission expired")), 302)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
