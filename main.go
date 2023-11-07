package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	dotenv "github.com/joho/godotenv"
	"go-web/routes"
	"net/http"
	"os"
)

func main() {
	loadEnv()
	router := chi.NewRouter()
	router.Use(getMiddlewares()...)
	routes.Routes(router)
	port := os.Getenv("SERVER_PORT")
	fmt.Printf("Server started at port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		fmt.Println(err)
		panic("Failed to start server.")
	}
}

func loadEnv() {
	if err := dotenv.Load(); err != nil {
		panic("Error loading .env file.")
	}
}

func getMiddlewares() []func(http.Handler) http.Handler {
	middlewares := []func(http.Handler) http.Handler{chiMiddleware.RealIP}
	debug := os.Getenv("DEBUG") == "true"
	if debug {
		middlewares = append(middlewares, chiMiddleware.Logger)
	}
	return middlewares
}
