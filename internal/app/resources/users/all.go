package users

import (

	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"intechno/httperrors"
	"net/http"
)

func (ur UserResource) All(w http.ResponseWriter, r *http.Request) error {
	idStr := r.Context().Value("_id").(string)
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		_ = render.Render(w, r, httperrors.Unauthorized(err))
		return id
	}

}