package policies

import (
	"fmt"
	"github.com/dealense7/documentSignatures/app/models"
	"github.com/dealense7/documentSignatures/exceptions"
	"github.com/dealense7/documentSignatures/initializers"
	"net/http"
	"strings"
)

func CheckPermission(model models.PermissionScope, permission string) *exceptions.Exception {
	if !initializers.AuthUser.HasPermission(model.GetScope() + "." + strings.ToLower(permission)) {
		message := fmt.Sprintf("you dont have %s permission for model %s", permission, model.GetScope())
		return exceptions.NewError(message, http.StatusUnauthorized)
	}

	return nil
}
