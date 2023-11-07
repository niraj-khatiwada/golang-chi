package contact

import (
	"github.com/go-chi/chi/v5"
	errors "go-web/libs"
	"html/template"
	"net/http"
)

func Contact(router chi.Router) {
	router.Get("/contact", func(w http.ResponseWriter, r *http.Request) {

		t, err := template.ParseFiles("views/contact.gohtml")
		if err != nil {
			errors.CatchRuntimeErrors(err)
			http.Error(w, "error", 500)
			return
		}

		if err := t.Execute(w, nil); err != nil {
			errors.CatchRuntimeErrors(err)
			http.Error(w, "error", 500)
		}
	})
}
