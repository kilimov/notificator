package drivers

import (
	"context"
	"github.com/kilimov/notificator/internal/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersRepository interface {
	Create(ctx context.Context, user *models.User) error
	All(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
