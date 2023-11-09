package routes

import (
	"github.com/go-chi/chi/v5"
	"go-web/routes/contact"
	"go-web/routes/root"
	"go-web/utils"
	"net/http"
)

func Routes(router chi.Router) {
	root.Root(router)
	contact.Contact(router)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		t := utils.ParseViewFiles(&w, "404.gohtml")
		if err := t.Execute(w, nil); err != nil {
			utils.CatchRuntimeErrors(err)
			http.Error(w, "error", 500)
			return
		}
	})
}
