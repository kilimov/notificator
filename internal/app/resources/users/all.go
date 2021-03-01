package users

import (
	"github.com/go-chi/render"
	"intechno/httperrors"
	"net/http"
)

func (ur UserResource) All(w http.ResponseWriter, r *http.Request) {
	users, err := ur.manager.All(r.Context())
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.JSON(w, r, users)
}
