package usecase

import (
	"fmt"

	auth "auth/domain"
	user1 "user/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthUsecase struct {
	AuthRepo auth.AuthRepository

	PasswordSrv auth.PasswordService
	TokenSrv    auth.TokenService
}

func NewAuthUsecase(authRepo auth.AuthRepository, passwordSrv auth.PasswordService, tokenSrv auth.TokenService) *AuthUsecase {
	return &AuthUsecase{
		AuthRepo: authRepo,

		PasswordSrv: passwordSrv,
		TokenSrv:    tokenSrv,
	}
}

func (u *AuthUsecase) RegisterUser(input user1.RegisterUser) (user1.User, error) {
	var user user1.User

	// Hash the user's password
	hashedPassword, err := u.PasswordSrv.HashPassword(input.Password)
	if err != nil {
		return user, err
	}

	// Create the user model
	user = user1.User{
		ID:            primitive.NewObjectID(),
		UserName:      input.UserName,
		Email:         input.Email,
		Password:      hashedPassword,
		Role:          input.Role,
		IsVerified:    false,
		OAuthProvider: "",
		OAuthID:       "",
	}

	err = u.AuthRepo.SaveUser(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *AuthUsecase) LoginUser(email, password string) (user1.User, string, string, error) {
	var user user1.User

	foundUser, err := u.AuthRepo.FindUserByEmail(email)
	if err != nil {
		return user, "", "", err
	}

	if foundUser == nil {
		return user, "", "", fmt.Errorf("user not found")
	}

	isMatch, err := u.PasswordSrv.ComparePassword(foundUser.Password, password)
	if err != nil {
		return user, "", "", err
	}

	if !isMatch {
		return user, "", "", fmt.Errorf("invalid password")
	}

	accessToken, err := u.TokenSrv.GenerateAccessToken(*foundUser)
	if err != nil {
		return user, "", "", err
	}

	refreshToken, err := u.TokenSrv.GenerateRefreshToken(*foundUser)
	if err != nil {
		return user, "", "", err
	}

	return *foundUser, accessToken, refreshToken, nil
}

func (u *AuthUsecase) RefreshTokens(refreshToken string) (string, string, error) {
	user, err := u.TokenSrv.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	newAccessToken, err := u.TokenSrv.GenerateAccessToken(*user)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := u.TokenSrv.GenerateRefreshToken(*user)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}
