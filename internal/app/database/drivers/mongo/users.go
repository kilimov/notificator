package mongo

import (
	"context"
	"github.com/kilimov/notificator/internal/app/database/drivers"
	"github.com/kilimov/notificator/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepo struct {
	collection *mongo.Collection
}

func (u UsersRepo) Create(ctx context.Context, user *models.User) error {
	if user == nil {
		return drivers.ErrUserDoesntExists
	}

	user.ID = primitive.NewObjectID()
	if _, err := u.collection.InsertOne(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u UsersRepo) All(ctx context.Context) ([]models.User, error) {
	users := make([]models.User, 0)

	cur, err := u.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (u UsersRepo) Update(ctx context.Context, user *models.User) error {
	if user == nil {
		return drivers.ErrUserDoesntExists
	}

	filter := bson.D{{"_id", user.ID}}
	update := bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "firstname", Value: user.FirstName},
			{Key: "lastname", Value: user.LastName},
			{Key: "email", Value: user.Email},
		},
	}}

	result, err := u.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return drivers.ErrNotModified
	}

	return nil
}

func (u UsersRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	_, err := u.collection.DeleteOne(ctx, filter)

	return err
}
