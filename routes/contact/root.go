package contact

import (
	"github.com/go-chi/chi/v5"
	"go-web/config"
	"go-web/utils"
	"net/http"
)

func Contact(router chi.Router, libs *config.Libs) {
	router.Get("/contact", func(w http.ResponseWriter, r *http.Request) {
		t := utils.ParseViewFiles(&w, "contact.gohtml")
		if err := t.Execute(w, nil); err != nil {
			utils.CatchRuntimeErrors(err)
			http.Error(w, "error", 500)
		}
	})
	router.Post("/contact", func(w http.ResponseWriter, r *http.Request) {
		// Req body
		// Validate it
		// Save to database
		t := utils.ParseViewFiles(&w, "contact.gohtml")
		if err := t.Execute(w, nil); err != nil {
			utils.CatchRuntimeErrors(err)
			http.Error(w, "error", 500)
		}
	})
}
