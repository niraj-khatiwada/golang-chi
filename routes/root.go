package routes

import (
	"github.com/go-chi/chi/v5"
	"go-web/config"
	"go-web/routes/contact"
	"go-web/routes/root"
	"go-web/utils"
	"net/http"
)

func Routes(router chi.Router, libs *config.Libs) {
	root.Root(router, libs)
	contact.Contact(router, libs)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		t := utils.ParseViewFiles(&w, "404.gohtml")
		if err := t.Execute(w, nil); err != nil {
			utils.CatchRuntimeErrors(err)
			http.Error(w, "error", 500)
			return
		}
	})
}
