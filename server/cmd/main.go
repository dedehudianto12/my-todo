package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	PORT := ":8080"

	fmt.Println("server is listening at ", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalln("server error", err)
	}
}