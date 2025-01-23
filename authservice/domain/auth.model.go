package domain

import (
	

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Enum-like type for User Roles
type Role string

const (
	Admin   Role = "admin"
	Creator Role = "creator"
	Viewer  Role = "viewer"
)

// UserClaims represents the claims for a user's JWT
type UserClaims struct {
	jwt.StandardClaims
	ID     primitive.ObjectID // User ID
	Avatar Media              // User avatar
	Name   string             // User name
	Email  string             // User email
	Role   string             // User role (e.g., admin, creator, viewer)
}


// EmailUserClaims represents the claims for an email verification JWT
type EmailUserClaims struct {
	jwt.StandardClaims
	ID    primitive.ObjectID `json:"_id"`   // User ID
	Email string             `json:"email"` // User email
}
type OAuthLoginUser struct {
	Provider string `json:"provider" bson:"provider"`
	Token    string `json:"token" bson:"token"`
}
