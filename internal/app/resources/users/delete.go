package users

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"intechno/httperrors"
	"net/http"
)

func (ur UserResource) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err := ur.manager.Delete(r.Context(), objID); err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
}
