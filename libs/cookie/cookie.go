package cookie

import (
	"go-web/utils"
	"net/http"
	"os"
	"strconv"
)

type Cookie struct {
	Name     string
	Value    string
	HttpOnly bool
	SameSite http.SameSite
	Secure   bool
	MaxAge   int
}

func (cookie *Cookie) ConstructCookie() http.Cookie {
	isHTTPS := os.Getenv("IS_HTTPS") == "true"
	cookie.Secure = isHTTPS
	if cookie.MaxAge == 0 {
		cookieExpiry := os.Getenv("JWT_EXPIRY_IN_MILLISECONDS")
		if expiryIn, err := strconv.ParseInt(cookieExpiry, 0, 64); err != nil {
			utils.CatchRuntimeErrors(err)
		} else {
			cookie.MaxAge = int(expiryIn)
		}
	}
	return http.Cookie{
		Name:     cookie.Name,
		Value:    cookie.Value,
		HttpOnly: cookie.HttpOnly,
		Secure:   cookie.Secure,
		SameSite: cookie.SameSite,
		MaxAge:   cookie.MaxAge,
	}

	//return fmt.Sprintf("%s=%s;MaxAge=%d;HttpOnly=%t;SameSite=%s;Secure=%t", cookie.Key, cookie.Value, cookie.ExpiresIn, cookie.HttpOnly, cookie.SameSite, cookie.Secure)
}
