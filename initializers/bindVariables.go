package initializers

import (
	"github.com/dealense7/documentSignatures/app/models"
)

var AuthUser models.User
var ApiVersion = 1

// ApiV1 supported api version
const ApiV1 int = 1

func SetUser(u models.User) {
	AuthUser = u
}

func SetApiVersion(version int) {
	ApiVersion = version
}
