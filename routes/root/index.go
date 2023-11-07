package root

import (
	"github.com/go-chi/chi/v5"
	errors "go-web/libs"
	"html/template"
	"net/http"
)

type Context struct {
	Name string
}

func Root(router chi.Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// If you want os support windows, provide path like this
		// filepath.Join("views", "root.gohtml") -> This is like path.join() in nodejs
		t, err := template.ParseFiles("views/root.gohtml")
		if err != nil {
			errors.CatchRuntimeErrors(err)
			http.Error(w, "error", 500)
			return
		}
		context := Context{Name: "World"}
		if err := t.Execute(w, context); err != nil {
			errors.CatchRuntimeErrors(err)
			http.Error(w, "error", 500)
			return
		}

	})
}
