package main

import (
	"go-web/libs/db"
	"os"
)

func main() {
	db.Migrate()
	os.Exit(0)
}
