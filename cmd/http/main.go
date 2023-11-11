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
	"path"
)

func main() {
	utils.LoadEnv()

	database, err := db.InitDB(config.Database{})
	if err != nil {
		log.Fatal("[error] Database connection error", err)
		return
	}

	libs := config.Libs{DB: database}

	router := chi.NewRouter()
	router.Use(utils.GetRouterMiddlewares()...)
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(path.Join(utils.GetCurrentDir(), "..", "..", "static")))))
	routes.Routes(router, &libs)

	port := os.Getenv("SERVER_PORT")
	fmt.Printf("Server started at port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		fmt.Println(err)
	}
}
