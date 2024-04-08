package policies

import (
	"fmt"
	"github.com/dealense7/documentSignatures/exceptions"
	"github.com/dealense7/documentSignatures/initializers"
	"github.com/dealense7/documentSignatures/models"
	"net/http"
	"strings"
)

func CheckPermission(model models.PermissionScope, permission string) *exceptions.Exception {
	permission = model.GetScope() + "." + strings.ToLower(permission)

	if !initializers.AuthUser.HasPermission(permission) {
		message := fmt.Sprintf("There is no %s permission for model %s", permission, model.GetScope())
		return exceptions.NewError(message, http.StatusUnauthorized)
	}

	return nil
}
