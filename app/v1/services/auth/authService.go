package authService

import (
	"github.com/dealense7/documentSignatures/app/models"
	"github.com/dealense7/documentSignatures/app/v1/repositories/user"
	"github.com/dealense7/documentSignatures/exceptions"
	"github.com/dealense7/documentSignatures/validation/requests/v1/auth"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func FindByEmail(email string) (*models.User, *exceptions.Exception) {
	item := userRepository.FindByEmail(email)

	if item == nil {
		return nil, exceptions.NewError("record not found", 404)
	}

	return item, nil
}

func Authorize(data authRequest.AuthRequest) (string, *exceptions.Exception) {

	item, userError := FindByEmail(data.Email)

	if userError != nil {
		return "", userError
	}

	err := bcrypt.CompareHashAndPassword([]byte(item.Password), []byte(data.Password))
	if err != nil {
		return "", exceptions.NewError("record not found", 404)
	}

	expireDate := time.Now().Add(30 * 24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": item.ID,
		"exp": expireDate,
	})

	tokenString, tokenErr := token.SignedString([]byte(os.Getenv("APP_KEY")))

	if tokenErr != nil {
		return "", exceptions.NewError("Failed to create token", 400)
	}

	return tokenString, nil
}
