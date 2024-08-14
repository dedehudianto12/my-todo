package usecase

import (
	"encoding/json"
	"log"
	"my-todo/internal/domain/models"
	"my-todo/internal/repository"
	"net/http"
)



type UserUsecase struct {
	userRepo *repository.UserRepository
}

func NewUserUsecase(userRepo *repository.UserRepository)*UserUsecase{
	return &UserUsecase{userRepo: userRepo}
}

func (uu *UserUsecase) CreateUser(w http.ResponseWriter, r *http.Request){
	var user *models.Register
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Error when decoding the body")
	}

	log.Println(user)
}