package business

import (
	"context"
	"github.com/kilimov/notificator/internal/app/database/drivers"
	"github.com/kilimov/notificator/internal/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserManager struct {
	repo drivers.UsersRepository
}

func NewUserManager(repo drivers.UsersRepository) *UserManager {
	return &UserManager{
		repo,
	}
}

func (um UserManager) Create(ctx context.Context, user *models.User) error {
	if user == nil {
		return UserCanNotBeEmpty
	}
	err := um.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (um UserManager) All(ctx context.Context) ([]models.User, error) {
	return um.repo.All(ctx)
}

func (um UserManager) Update(ctx context.Context, user *models.User) error {
	if user == nil {
		return UserCanNotBeEmpty
	}
	err := um.repo.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (um UserManager) Delete(ctx context.Context, id primitive.ObjectID) error {
	return um.repo.Delete(ctx, id)
}
