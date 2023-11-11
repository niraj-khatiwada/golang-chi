package main

import (
	"bufio"
	"fmt"
	"go-web/libs/db"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Are you sure you want to run migration? (yes/no)")
	reader := bufio.NewReader(os.Stdin)
	option, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("[error] error executing migration script ", err)
	}
	if strings.TrimSpace(option) == "yes" {
		db.Migrate()
	} else {
		fmt.Println("Migration halted.")
	}
	os.Exit(0)
}
