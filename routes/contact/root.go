package routes

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	validator "github.com/go-ozzo/ozzo-validation/v4"
	validatorIs "github.com/go-ozzo/ozzo-validation/v4/is"
	"go-web/libs"
	"go-web/libs/cookie"
	"go-web/libs/redis"
	"go-web/libs/validation"
	"go-web/models"
	"go-web/utils"
	"go-web/views"
	"net/http"
	"time"
)

type ContactForm struct {
	Name        string
	Email       string
	Description string
}

type ContactData struct {
	Form             ContactForm
	ValidationErrors validation.ValidationErrors
	Success          bool
	Message          string
}

func Contact(rootRouter chi.Router, libs *libs.Libs) {
	rootRouter.Route("/contact", func(router chi.Router) {
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			UUID := utils.AlphanumericUUID{}

			if uuid, err := UUID.Generate(); err != nil {
				utils.CatchRuntimeErrors(err)
				http.Error(w, "error", 500)
				return
			} else {
				if redisErr := libs.Redis.Set(fmt.Sprintf("csrf-token-%s", uuid), uuid, 15*time.Minute); redisErr != nil {
					utils.CatchRuntimeErrors(redisErr)
					http.Error(w, "error", 500)
					return
				}
				cookieOptions := &cookie.Cookie{HttpOnly: true, Name: "csrf", Value: uuid}
				c := cookieOptions.ConstructCookie()
				http.SetCookie(w, &c)
			}

			t := views.ParseFiles(&w, "contact.gohtml")
			if err := t.Execute(w, nil); err != nil {
				utils.CatchRuntimeErrors(err)
				http.Error(w, "error", 500)
			}
		})
		router.Group(func(router chi.Router) {
			nameRules := []validator.Rule{validator.Required, validator.Length(2, 100)}
			emailRules := []validator.Rule{validator.Required, validatorIs.Email}
			descriptionRules := []validator.Rule{validator.Required, validator.Length(1, 1000)}
			validations := map[string][]validator.Rule{"name": nameRules, "email": emailRules, "description": descriptionRules}
			validationMiddleware := validation.WithRequestInputValidation{}
			router.Use(validationMiddleware.Validate(validations))
			router.Post("/", func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context().Value("WithRequestInputValidation").(validation.WithRequestInputValidationContext)
				formData := ctx.FormData
				data := ContactData{}
				data.Form = ContactForm{Name: formData.Get("name"), Email: formData.Get("email"), Description: formData.Get("description")}
				var validationErrors = ctx.ValidationErrors
				template := views.ParseFiles(&w, "contact.gohtml")
				if validationErrors != nil {
					data.ValidationErrors = validationErrors
					if err := template.Execute(w, data); err != nil {
						utils.CatchRuntimeErrors(err)
						http.Error(w, "error", 500)
					}
					return
				}

				// CSRF
				cookieStr, cookieErr := r.Cookie("csrf")
				if cookieErr != nil {
					if errors.Is(cookieErr, http.ErrNoCookie) {
						// TODO: Integrate Session and redirect back to /contact
						data.Success = false
						data.Message = "Invalid form submission. Form submission time expired."
						data.Form = ContactForm{}
						if err := template.Execute(w, data); err != nil {
							utils.CatchRuntimeErrors(err)
							http.Error(w, "error", 500)
						}
					}
					if err := template.Execute(w, data); err != nil {
						utils.CatchRuntimeErrors(err)
						http.Error(w, "error", 500)
					}
					return
				}
				var csrfToken string
				csrfTokenKey := fmt.Sprintf("csrf-token-%s", cookieStr.Value)
				redisErr := libs.Redis.Get(csrfTokenKey, &csrfToken)
				if redisErr != nil {
					if errors.Is(redisErr, redis.ErrNoEntry) {
						// TODO: Integrate Session and redirect back to /contact
						data.Success = false
						data.Message = "Invalid form submission. Form submission time expired."
						data.Form = ContactForm{}
					}
					if err := template.Execute(w, data); err != nil {
						utils.CatchRuntimeErrors(err)
						http.Error(w, "error", 500)
					}
					return
				}

				result := libs.DB.Client.Create(&models.Contact{Name: data.Form.Name, Email: data.Form.Email, Description: data.Form.Description})
				if result.Error != nil {
					utils.CatchRuntimeErrors(result.Error)
					http.Error(w, "error", 500)
					return
				}
				// Delete CSRF Token
				if redisDelErr := libs.Redis.Delete([]string{csrfTokenKey}...); redisDelErr != nil {
					utils.CatchRuntimeErrors(result.Error)
					http.Error(w, "error", 500)
					return
				}
				// Instead of doing this, we can just redirect to /contact path. But for that we need some kind of session to flash the message.
				// TODO: Integrate Session
				data.Success = true
				data.Message = "Thank you for contacting. We'll get back to you soon."
				data.Form = ContactForm{}
				if err := template.Execute(w, data); err != nil {
					utils.CatchRuntimeErrors(err)
					http.Error(w, "error", 500)
				}
			})
		})
	})

}
