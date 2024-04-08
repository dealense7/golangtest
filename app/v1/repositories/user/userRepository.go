package userRepository

import (
	"errors"
	"github.com/dealense7/documentSignatures/app/models"
	"github.com/dealense7/documentSignatures/initializers"
)

func FindItems() []models.User {
	var items []models.User
	initializers.DB.Find(&items)

	return items
}

func FindByEmail(email string) *models.User {
	var user models.User
	initializers.DB.First(&user, "email = ?", email)

	if user.ID != 0 {
		return &user
	}

	return nil
}

func FindByEmailUnscoped(email string) *models.User {
	var user models.User
	initializers.DB.Unscoped().First(&user, "email = ?", email)

	if user.ID != 0 {
		return &user
	}

	return nil
}

func FindById(id int) *models.User {
	var user models.User
	initializers.DB.First(&user, "id = ?", id)

	if user.ID != 0 {
		return &user
	}

	return nil
}

func Create(item *models.User) error {
	result := initializers.DB.Create(item)

	if result.Error != nil {
		return errors.New("record is not added to database")
	}

	return nil
}

func Update(item *models.User, data models.User) error {

	result := initializers.DB.Model(&item).Updates(data)

	if result.Error != nil {
		return errors.New("record is not updated")
	}

	return nil
}

func Delete(item *models.User) {
	initializers.DB.Delete(&item, item.ID)
}
