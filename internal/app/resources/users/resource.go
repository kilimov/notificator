package users

import (
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserResource struct {
	collection *mongo.Collection
}

func (ur UserResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {

		r.Post("/users", ur.Create)
		//r.Get("/users", ur.All)
		//r.Put("/users", ur.Update)
		r.Delete("/users", ur.Delete)
	})

	return r
}
.