package routes

import (
	"github.com/go-chi/chi/v5"
	"go-web/routes/contact"
	"go-web/routes/root"
	"html/template"
	"net/http"
)

func Routes(router chi.Router) {
	root.Root(router)
	contact.Contact(router)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("views/404.gohtml")
		if err != nil {
			return
		}
		if err2 := t.Execute(w, nil); err2 != nil {
			return
		}
	})
}
