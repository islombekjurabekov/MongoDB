package services

import (
	"MongoDB/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImplement struct {
	UserCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) *UserServiceImplement {
	return &UserServiceImplement{
		UserCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImplement) CreateUser(user *models.User) error {
	_, err := u.UserCollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImplement) GetUser(name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.UserCollection.FindOne(u.ctx, query).Decode(&user)

	return user, err
}

func (u *UserServiceImplement) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.UserCollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("Nothing not found")
	}

	return users, nil
}

func (u *UserServiceImplement) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_age", Value: user.Age}, bson.E{Key: "user_address", Value: user.Address}}}}
	result, _ := u.UserCollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("The data not found for update")
	}
	return nil
}

func (u *UserServiceImplement) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	result, _ := u.UserCollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("The data not found for delete")
	}
	return nil
}
