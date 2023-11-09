package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"go-web/config"
	"go-web/libs/db"
	"go-web/routes"
	"go-web/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	utils.LoadEnv()

	_, dbErr := db.InitializeDB(config.Database{})
	if dbErr != nil {
		log.Fatal("Database connection error", dbErr)
		return
	}

	router := chi.NewRouter()
	router.Use(utils.GetRouterMiddlewares()...)
	routes.Routes(router)

	port := os.Getenv("SERVER_PORT")
	fmt.Printf("Server started at port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		fmt.Println(err)
		panic("Failed to start server.")
	}
}
