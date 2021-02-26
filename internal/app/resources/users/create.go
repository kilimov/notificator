package users

import (
	"github.com/kilimov/notificator/internal/app/models"
	"net/http"
)

type AddUserRequest struct {
	user models.User `json:"sku"`
}


func (ur UserResource) Create(w http.ResponseWriter, r *http.Request) {



	return
}