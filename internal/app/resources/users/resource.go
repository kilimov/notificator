package users

import (
	"github.com/go-chi/chi"
)

type UserResource struct {
}

func (ur UserResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {

		//r.Post("/users", ur.Create)
		//r.Get("/users", ur.Read)
		//r.Put("/users", ur.Update)
		//r.Delete("/users", ur.Delete)
	})

	return r
}
