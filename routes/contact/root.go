package contact

import (
	"github.com/go-chi/chi/v5"
	validator "github.com/go-ozzo/ozzo-validation/v4"
	validatorIs "github.com/go-ozzo/ozzo-validation/v4/is"
	"go-web/config"
	"go-web/libs/validation"
	"go-web/models"
	"go-web/utils"
	"go-web/views"
	"net/http"
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

func Contact(rootRouter chi.Router, libs *config.Libs) {
	rootRouter.Route("/contact", func(router chi.Router) {
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
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

				result := libs.DB.Create(&models.Contact{Name: data.Form.Name, Email: data.Form.Email, Description: data.Form.Description})
				if result.Error != nil {
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
