package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"reflect"
)

func main() {
	router := chi.NewRouter()
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.RealIP)
	handler(router)
	fmt.Println("Server started at port 3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		panic("Failed to start server.")
	}
}

func handler(router chi.Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("<h1>Hello World</h1>"))
		if err != nil {
			log.Fatal("Error writing.")
		}
	})
	// Named Param /{id}. Jus like /:id in Express
	// Query. /?abc=def&abc=def -> abc=[def ghi]

	router.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		fmt.Println(query["abc"], reflect.TypeOf(query["abc"]))
		userId := chi.URLParam(r, "id")
		_, err := w.Write([]byte(fmt.Sprintf("<h1>Hello User %s</h1>", userId)))
		if err != nil {
			log.Fatal("Error writing.")
		}
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("<h1>Page not found</h1>")); err != nil {
		}
	})
}
