package route

import (
	"my-todo/internal/usecase"

	"github.com/gorilla/mux"
)

func NewRouter(uu *usecase.UserUsecase) *mux.Router{
	r := mux.NewRouter()

	r.HandleFunc("/register", uu.CreateUser).Methods("POST")

	return r
}