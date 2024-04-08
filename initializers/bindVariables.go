package initializers

import (
	"github.com/dealense7/documentSignatures/models"
)

var AuthUser models.User

func SetUser(u models.User) {
	AuthUser = u
}
