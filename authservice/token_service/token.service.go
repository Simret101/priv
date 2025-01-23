package token_service

import (
	"errors"
	"time"

	auth1 "auth/domain"
	user1 "user/domain"

	"github.com/dgrijalva/jwt-go"
)

type TokenService_imp struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
}

func NewTokenService(accessSecret, refreshSecret string) *TokenService_imp {
	return &TokenService_imp{
		AccessTokenSecret:  accessSecret,
		RefreshTokenSecret: refreshSecret,
	}
}

func (t *TokenService_imp) GenerateAccessToken(user user1.User) (string, error) {
	claims := auth1.UserClaims{
		ID:     user.ID,
		Name:   user.UserName,
		Avatar: auth1.Media(user.ProfilePicture),
		Email:  user.Email,
		Role:   string(user.Role), // Assign user role
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.AccessTokenSecret))
}

func (t *TokenService_imp) GenerateRefreshToken(user user1.User) (string, error) {
	claims := auth1.UserClaims{
		ID:     user.ID,
		Name:   user.UserName,
		Avatar: auth1.Media(user.ProfilePicture),
		Email:  user.Email,
		Role:   string(user.Role), // Assign user role
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.RefreshTokenSecret))
}

func (t *TokenService_imp) ValidateAccessToken(tokenStr string) (*user1.User, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &auth1.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.AccessTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid access token")
	}

	claims, ok := token.Claims.(*auth1.UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return &user1.User{
		ID:             claims.ID,
		UserName:       claims.Name,
		ProfilePicture: user1.Media(claims.Avatar),
		Email:          claims.Email,
		// Return role instead of IsAdmin
	}, nil
}

func (t *TokenService_imp) ValidateRefreshToken(tokenStr string) (*user1.User, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &auth1.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.RefreshTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(*auth1.UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return &user1.User{
		ID:             claims.ID,
		UserName:       claims.Name,
		ProfilePicture: user1.Media(claims.Avatar),
		Email:          claims.Email,
		// Return role instead of IsAdmin
	}, nil
}
