package repository

import (
	"context"
	"user/database"
	"user/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	Collection database.CollectionInterface
}

func NewUserRepository(collection database.CollectionInterface) *UserRepository {
	return &UserRepository{Collection: collection}
}

func (repository *UserRepository) GetUserDocumentByID(id string) (domain.User, error) {
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}
	filter := bson.D{{Key: "_id", Value: pid}}
	result := repository.Collection.FindOne(context.TODO(), filter)

	var user domain.User
	err = result.Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repository *UserRepository) GetUserDocuments() ([]domain.User, error) {
	filter := bson.D{{}}
	cursor, err := repository.Collection.Find(context.TODO(), filter)

	if err != nil {
		return []domain.User{}, err
	}

	users := []domain.User{}
	for cursor.Next(context.TODO()) {
		var user domain.User
		err := cursor.Decode(&user)
		if err != nil {
			return []domain.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository *UserRepository) UpdateUserDocument(id string, user domain.UpdateUser) (domain.ResponseUser, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.Collection.UpdateOne(context.TODO(), bson.M{"_id": obId}, bson.M{"$set": user})
	return domain.ResponseUser{}, err
}

func (repository *UserRepository) UpdateUserPassword(id string, new_hashed_password string) (domain.User, error) {
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}

	filter := bson.D{{Key: "_id", Value: pid}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: new_hashed_password}}}}

	result := repository.Collection.FindOneAndUpdate(context.TODO(), filter, update)
	var user domain.User

	err = result.Decode(&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepository) DeleteUserDocument(id string) error {
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: pid}}
	_, err = repository.Collection.DeleteOne(context.TODO(), filter)
	return err
}

func (repository *UserRepository) FilterUserDocument(filter map[string]string) ([]domain.User, error) {
	filters := bson.D{}
	for key, value := range filter {
		temp := bson.E{Key: key, Value: value}
		filters = append(filters, temp)
	}

	cursor, err := repository.Collection.Find(context.TODO(), filters)
	if err != nil {
		return []domain.User{}, err
	}
	users := []domain.User{}
	for cursor.Next(context.TODO()) {
		user := domain.User{}
		err := cursor.Decode(&user)
		if err != nil {
			return []domain.User{}, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository *UserRepository) PromoteUser(id string) (domain.User, error) {
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}

	filter := bson.D{{Key: "_id", Value: pid}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "role", Value: "admin"}}}}

	updated_user := repository.Collection.FindOneAndUpdate(context.TODO(), filter, update)
	var new_user domain.User
	err = updated_user.Decode(&new_user)
	if err != nil {
		return domain.User{}, err
	}

	return new_user, nil
}

func (repository *UserRepository) DemoteUser(id string) (domain.User, error) {
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}

	filter := bson.D{{Key: "_id", Value: pid}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "role", Value: "viewer"}}}}

	updated_user := repository.Collection.FindOneAndUpdate(context.TODO(), filter, update)
	var new_user domain.User
	err = updated_user.Decode(&new_user)
	if err != nil {
		return domain.User{}, err
	}

	return new_user, nil
}
