package seeds

import (
	"github.com/dealense7/documentSignatures/app/models"
	"github.com/dealense7/documentSignatures/initializers"
	"github.com/jaswdr/faker/v2"
	"golang.org/x/crypto/bcrypt"
)

func UserSeed() {
	numUsers := 10
	fake := faker.New()

	password, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)

	admin := models.User{
		FistName: fake.Person().Name(),
		LastName: fake.Person().LastName(),
		Email:    models.AdminEmail,
		Password: string(password),
	}

	initializers.DB.Where(models.User{Email: admin.Email}).FirstOrCreate(&admin)

	for i := 0; i < numUsers; i++ {
		user := models.User{
			FistName: fake.Person().Name(),
			LastName: fake.Person().LastName(),
			Email:    fake.Internet().Email(),
			Password: string(password),
		}
		initializers.DB.Where(models.User{Email: user.Email}).FirstOrCreate(&user)
	}

}
