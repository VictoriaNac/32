package service

import (
	user_app "github.com/VictoriaNac/finish"
	"finish/server_app/interial/repository"
)

type User interface {
	CreateUser(user user_app.RequestCreate) (string, error)
	MakeFriends(sourceId, targetId string) (string, error)
	DeleteUser(id string) (string, error)
	GetFriends(id string) ([]string, error)
	UpdateAge(id, age string) (string, error)
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}