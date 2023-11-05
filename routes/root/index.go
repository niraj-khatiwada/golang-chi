package root

import (
	"fmt"
	"github.com/alecthomas/template"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Context struct {
	Name string
}

func Root(router chi.Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("views/root.gohtml")
		if err != nil {
			fmt.Println("Error parsing file.")
			return
		}
		context := Context{Name: "Niraj"}
		if err2 := t.Execute(w, context); err2 != nil {
			return
		}
	})
}
