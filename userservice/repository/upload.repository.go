package repository

import (
	"context"
	"errors"
	"user/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UploadRepo struct {
	UserRepository
}

func NewUploadRepository(user_repo UserRepository) *UploadRepo {
	return &UploadRepo{
		UserRepository: user_repo,
	}
}

func (repo *UploadRepo) AddProfile(media domain.Media, id string) error {
	// Generate ObjectID from user ID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID")
	}
	media.ID = primitive.NewObjectID()

	// Create the MongoDB filter and update
	filter := bson.D{{Key: "_id", Value: objID}}
	data := bson.D{{Key: "profile_picture", Value: media}}
	setter := bson.D{{Key: "$set", Value: data}}

	// Update the user document with the new profile picture
	_, err = repo.UserRepository.Collection.UpdateOne(context.TODO(), filter, setter)
	if err != nil {
		return err
	}

	// Successfully updated the profile picture
	return nil
}
