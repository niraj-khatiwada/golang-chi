package root

import (
	"github.com/go-chi/chi/v5"
	"go-web/config"
	"go-web/utils"
	"go-web/views"
	"net/http"
)

type Metadata struct {
	Title string
}
type Context struct {
	Name     string
	Metadata Metadata
}

func Root(router chi.Router, _ *config.Libs) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t := views.ParseFiles(&w, "root.gohtml")
		context := Context{Name: "World", Metadata: Metadata{Title: "Golang"}}
		if err := t.Execute(w, context); err != nil {
			utils.CatchRuntimeErrors(err)
			http.Error(w, "error", 500)
			return
		}
	})
}
