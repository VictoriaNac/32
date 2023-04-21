package hendler

import (
	"github.com/go-chi/chi/v5"
	"finish/server_app/interial/service"
	
)

type Hendler struct {
	services *service.Service
}

func NewHendler(services *service.Service) *Hendler {
	return &Hendler{services: services}
}

func (h *Hendler) InitRouters() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/create", h.CreateUser)
	router.Post("/make_friends", h.MakeFriends)
	router.Delete("/user", h.DeleteUser)
	router.Get("/friends/{user_id}", h.GetFriends)
	router.Put("/{user_id}", h.UpdateAge)

	return router
}