package users

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/kilimov/notificator/internal/app/models"
	"intechno/httperrors"
	"net/http"
)

func (ur UserResource) Create(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err := ur.manager.Create(r.Context(), user); err != nil {
		_ = render.Render(w, r, httperrors.UnprocessableEntity(err))
		return
	}

	w.WriteHeader(http.StatusCreated)
}
