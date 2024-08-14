package main

import (
	"fmt"
	"log"
	"my-todo/internal/delivery/route"
	"my-todo/internal/repository"
	"my-todo/internal/usecase"
	"my-todo/pkg/db"
	"net/http"
)

func main() {

	db := db.InitDB()
	defer db.Close()

	repo := repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(&repo)
	route := route.NewRouter(usecase)

	PORT := ":8080"

	fmt.Println("server is listening at ", PORT)
	err := http.ListenAndServe(PORT, route)
	if err != nil {
		log.Fatalln("server error", err)
	}
}