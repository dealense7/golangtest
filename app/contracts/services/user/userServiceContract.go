package userServiceContract

import (
	"github.com/dealense7/documentSignatures/app/models"
	"github.com/dealense7/documentSignatures/app/v1/services/user"
	"github.com/dealense7/documentSignatures/exceptions"
	"github.com/dealense7/documentSignatures/initializers"
	"github.com/dealense7/documentSignatures/validation/requests/v1/auth"
	"github.com/dealense7/documentSignatures/validation/requests/v1/user"
)

type UserServiceContract interface {
	FindItems() ([]models.User, *exceptions.Exception)
	FindById(id int) *models.User
	FindByIdOrFail(id int) (*models.User, *exceptions.Exception)
	Create(data authRequest.UserRegistrationRequest) (*models.User, *exceptions.Exception)
	Update(item *models.User, data userRequest.UpdateUserRequest) *exceptions.Exception
	Delete(item *models.User)
}

func NewUserService() UserServiceContract {
	version := initializers.ApiVersion

	switch version {
	case initializers.ApiV1:
		return userService.NewUserService()
	default:
		return userService.NewUserService()
	}
}
