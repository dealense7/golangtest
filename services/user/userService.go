package userService

import (
	"github.com/dealense7/documentSignatures/exceptions"
	"github.com/dealense7/documentSignatures/models"
	"github.com/dealense7/documentSignatures/policies"
	userRepository "github.com/dealense7/documentSignatures/repositories/user"
	authRequest "github.com/dealense7/documentSignatures/requests/auth"
	userRequest "github.com/dealense7/documentSignatures/requests/user"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func FindItems() ([]models.User, *exceptions.Exception) {
	err := policies.CheckPermission(models.User{}, "View")

	if err != nil {
		return nil, err
	}

	return userRepository.FindItems(), nil
}

func FindByEmail(email string) (*models.User, *exceptions.Exception) {
	item := userRepository.FindByEmail(email)

	if item == nil {
		return nil, exceptions.NewError("record not found", 404)
	}

	return item, nil
}

func FindById(id int) *models.User {
	return userRepository.FindById(id)
}

func FindByIdOrFail(id int) (*models.User, *exceptions.Exception) {
	item := FindById(id)

	if item == nil {
		return nil, exceptions.NewError("record not found", 404)
	}

	return item, nil
}

func Create(data authRequest.UserRegistrationRequest) (*models.User, *exceptions.Exception) {
	// Check if email is used
	user := userRepository.FindByEmailUnscoped(data.Email)
	if user != nil {
		return nil, exceptions.NewError("email is already used", 422)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return nil, exceptions.NewError("problem with hashing password", 400)
	}

	item := models.User{
		FistName: data.FirstName,
		LastName: data.LastName,
		Email:    data.Email,
		Password: string(hashedPassword),
	}

	err = userRepository.Create(&item)

	if err != nil {
		return nil, exceptions.NewError(err.Error(), 404)
	}

	return &item, nil
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

func Update(item *models.User, data userRequest.UpdateUserRequest) *exceptions.Exception {

	dto := models.User{FistName: data.FirstName, LastName: data.LastName, Email: data.Email}

	err := userRepository.Update(item, dto)

	if err != nil {
		return exceptions.NewError(err.Error(), 400)
	}

	return nil
}

func Delete(item *models.User) {
	userRepository.Delete(item)
}
