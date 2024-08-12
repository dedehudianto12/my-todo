package main

import (
	"fmt"
	"log"
	"my-todo/pkg/db"
	"net/http"
)

func main() {

	db := db.InitDB()
	defer db.Close()

	PORT := ":8080"

	fmt.Println("server is listening at ", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalln("server error", err)
	}
}