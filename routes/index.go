package routes

import (
	"github.com/alecthomas/template"
	"github.com/go-chi/chi/v5"
	"go-web/routes/root"
	"net/http"
)

func Routes(router chi.Router) {
	root.Root(router)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("views/404.gohtml")
		if err != nil {
			return
		}
		var context interface{}
		if err2 := t.Execute(w, context); err2 != nil {
			return
		}
	})
}
