package users

import (
	"github.com/go-chi/chi"
	"github.com/kilimov/notificator/internal/app/business"
)

type UserResource struct {
	manager *business.UserManager
}

func NewUserResource(manager *business.UserManager) *UserResource {
	return &UserResource{
		manager: manager,
	}
}

func (ur UserResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {

		r.Post("/", ur.Create)
		r.Get("/", ur.All)
		//r.Put("/", ur.Update)
		r.Delete("/{id}", ur.Delete)
	})

	return r
}
