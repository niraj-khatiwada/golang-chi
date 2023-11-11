package views

import (
	"embed"
	"go-web/utils"
	"html/template"
	"net/http"
)

//go:embed *
var FS embed.FS

func ParseFiles(writer *http.ResponseWriter, filenames ...string) *template.Template {
	filenames = append(filenames, "partials/*.gohtml")
	t, err := template.ParseFS(FS, filenames...)
	if err != nil {
		utils.CatchRuntimeErrors(err)
		http.Error(*writer, "error", 500)
		return nil
	}
	return t
}
