package utils

import (
	"go-web/views"
	"html/template"
	"net/http"
)

func ParseViewFiles(writer *http.ResponseWriter, filenames ...string) *template.Template {
	t, err := template.ParseFS(views.FS, filenames...)
	if err != nil {
		CatchRuntimeErrors(err)
		http.Error(*writer, "error", 500)
		return nil
	}
	return t
}
