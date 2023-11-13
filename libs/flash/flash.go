package flash

import (
	"fmt"
	"go-web/libs/base64"
	"net/http"
	"time"
)

type Flash struct{}

const (
	FlashSuccess = "FlashSuccess"
	FlashError   = "FlashError"
)

func (flash *Flash) GetFlashMessage(w http.ResponseWriter, r *http.Request, name string, value *string) error {
	key := flash.generateKey(name)
	cookie, err := r.Cookie(key)
	if err != nil {
		return err
	}
	var target string
	if decodeErr := base64.Decode(cookie.Value, &target); decodeErr != nil {
		return decodeErr
	}
	*value = target
	setCookie := &http.Cookie{Name: key, MaxAge: -1, Expires: time.Unix(100, 0)}
	http.SetCookie(w, setCookie)
	return nil
}

func (flash *Flash) SetFlashMessage(w http.ResponseWriter, name string, value string) {
	key := flash.generateKey(name)
	encodedValue := base64.Encode(value)
	http.SetCookie(w, &http.Cookie{Name: key, Value: encodedValue})
}

func (flash *Flash) generateKey(key string) string {
	return fmt.Sprintf("flash-%s", key)
}
