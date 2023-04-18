package repository

import  
user_app "finish"

type User interface {
	CreateUser(user user_app.RequestCreate) (string, error)
	MakeFriends(sourceId, targetId string) (string, error)
	DeleteUser(id string) (string, error)
	GetFriends(id string) ([]string, error)
	UpdateAge(id, age string) (string, error)
	getUser(id string) (*UserItem, error)
	setUser(id string, user *UserItem) error
}

type Repository struct {
	User
}

func NewRepository(db *DataBase) *Repository {
	return &Repository{
		User: NewUserDB(db),
	}
}


func sliceToMap(slice []string) (map[string]bool, error) {
	dict := make(map[string]bool, len(slice))

	for _, item := range slice {
		dict[item] = true
	}

	return dict, nil
}


func mapToSlice(dict map[string]bool) ([]string, error) {
	var slice []string

	for keys := range dict {
		slice = append(slice, keys)
	}

	return slice, nil
}