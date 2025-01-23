package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Enum-like type for User Roles
type Role string

const (
	Admin   Role = "admin"
	Creator Role = "creator"
	Viewer  Role = "viewer"
)

// Actual user model
type User struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName       string             `json:"username" bson:"username"`
	Bio            string             `json:"bio,omitempty" bson:"bio,omitempty"`
	ProfilePicture Media              `json:"profile_picture,omitempty" bson:"profile_picture,omitempty"`
	Email          string             `json:"email" bson:"email"`
	Role           Role               `json:"role" bson:"role"` // Replacing Is_Admin with Role field
	Password       string             `json:"password,omitempty" bson:"password,omitempty"`
	IsVerified     bool               `json:"is_verified" bson:"is_verified"`
	OAuthProvider  string             `json:"oauth_provider,omitempty" bson:"oauth_provider,omitempty"`
	OAuthID        string             `json:"oauth_id,omitempty" bson:"oauth_id,omitempty"`
}

// Response model for User that will be returned from the server
type ResponseUser struct {
	ID             string `json:"_id" bson:"_id"`
	UserName       string `json:"username" bson:"username"`
	Bio            string `json:"bio,omitempty" bson:"bio,omitempty"`
	ProfilePicture Media  `json:"profile_picture,omitempty" bson:"profile_picture,omitempty"`
	Email          string `json:"email" bson:"email"`
	Role           Role   `json:"role" bson:"role"` // Updating the response model to use Role
}

type UpdateUser struct {
	UserName   string `json:"username" bson:"username"`
	Bio        string `json:"bio,omitempty" bson:"bio,omitempty"`
	IsVerified bool   `json:"is_verified" bson:"is_verified"`
	Role       Role   `json:"role" bson:"role"` // Adding role to the update model
}

type LogINUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type OAuthLoginUser struct {
	Provider string `json:"provider" bson:"provider"`
	Token    string `json:"token" bson:"token"`
}

type RegisterUser struct {
	UserName string `json:"username" bson:"username" validate:"required,min=3,max=30,alphanum"`
	Bio      string `json:"bio,omitempty" bson:"bio,omitempty"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password,omitempty" bson:"password,omitempty" validate:"required,min=8"`
	Role     Role   `json:"role" bson:"role"` // Adding role to the register model
}

type UpdatePassword struct {
	Password        string `json:"password" bson:"password"`
	ConfirmPassword string `json:"confirm_password" bson:"confirm_password"`
}

type VerifyEmail struct {
	Email string `json:"email" bson:"email"`
}

// Convert the actual User model to the ResponseUser model
func CreateResponseUser(user User) ResponseUser {
	return ResponseUser{
		ID:             user.ID.Hex(),
		UserName:       user.UserName,
		Bio:            user.Bio,
		ProfilePicture: user.ProfilePicture,
		Email:          user.Email,
		Role:           user.Role, // Using the updated role field
	}
}
