package root

import (
	"github.com/go-chi/chi/v5"
	"go-web/utils"
	"net/http"
)

type Context struct {
	Name string
}

func Root(router chi.Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t := utils.ParseViewFiles(&w, "root.gohtml")
		context := Context{Name: "World"}
		if err := t.Execute(w, context); err != nil {
			utils.CatchRuntimeErrors(err)
			http.Error(w, "error", 500)
			return
		}

	})
}
