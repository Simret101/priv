package repository

import (
	"context"
	"errors"
	"regexp"

	"auth/database"
	user "user/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthRepo struct {
	Collection database.CollectionInterface
}

func (repo *AuthRepo) EnsureIndexes() error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := repo.Collection.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}

func NewAuthRepo(coll database.CollectionInterface) (*AuthRepo, error) {
	AR := &AuthRepo{
		Collection: coll,
	}

	// Ensure indexes are created
	if err := AR.EnsureIndexes(); err != nil {
		return nil, err
	}

	return AR, nil
}

func (repo *AuthRepo) SaveUser(user *user.User) error {
	// Validate the user before saving
	if err := validateUser(user); err != nil {
		return err
	}

	_, err := repo.Collection.InsertOne(context.TODO(), user)
	return err
}

func (repo *AuthRepo) FindUserByEmail(email string) (*user.User, error) {
	var user user.User
	err := repo.Collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}


// validateUser validates the user struct before saving it to the database
func validateUser(user *user.User) error {
	// Ensure email is not empty and matches a basic email regex
	if user.Email == "" {
		return errors.New("email is required")
	}
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, user.Email)
	if !matched {
		return errors.New("invalid email format")
	}

	// Ensure password is not empty
	if user.Password == "" {
		return errors.New("password is required")
	}

	// Ensure username is not empty
	if user.UserName == "" {
		return errors.New("username is required")
	}

	// Ensure role is valid
	validRoles := map[string]bool{
		"admin":   true,
		"creator": true,
		"viewer":  true,
	}
	if _, isValid := validRoles[string(user.Role)]; !isValid {
		return errors.New("invalid role")
	}

	return nil
}
