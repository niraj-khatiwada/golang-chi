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
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.RealIP)
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
