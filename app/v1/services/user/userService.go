package userService

import (
	"github.com/dealense7/documentSignatures/app/models"
	"github.com/dealense7/documentSignatures/app/v1/repositories/user"
	"github.com/dealense7/documentSignatures/exceptions"
	"github.com/dealense7/documentSignatures/validation/policies"
	"github.com/dealense7/documentSignatures/validation/requests/v1/auth"
	"github.com/dealense7/documentSignatures/validation/requests/v1/user"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) FindItems() ([]models.User, *exceptions.Exception) {
	err := policies.CheckPermission(models.User{}, "View")

	if err != nil {
		return nil, err
	}

	return userRepository.FindItems(), nil
}

func (s *UserService) FindById(id int) *models.User {
	return userRepository.FindById(id)
}

func (s *UserService) FindByIdOrFail(id int) (*models.User, *exceptions.Exception) {
	item := s.FindById(id)

	if item == nil {
		return nil, exceptions.NewError("record not found", 404)
	}

	err := policies.CheckPermission(item, "View")

	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *UserService) Create(data authRequest.UserRegistrationRequest) (*models.User, *exceptions.Exception) {
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

func (s *UserService) Update(item *models.User, data userRequest.UpdateUserRequest) *exceptions.Exception {
	// Check if email is used
	userWithSameMail := userRepository.FindByEmailUnscoped(data.Email)
	if userWithSameMail != nil && userWithSameMail.ID != item.ID {
		return exceptions.NewError("email is already used", 422)
	}

	dto := models.User{FistName: data.FirstName, LastName: data.LastName, Email: data.Email}

	err := userRepository.Update(item, dto)

	if err != nil {
		return exceptions.NewError(err.Error(), 400)
	}

	return nil
}

func (s *UserService) Delete(item *models.User) {
	userRepository.Delete(item)
}
