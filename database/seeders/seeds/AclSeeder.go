package seeds

import (
	"github.com/dealense7/documentSignatures/initializers"
	"github.com/dealense7/documentSignatures/models"
)

func AclSeed() {

	type Acl struct {
		title       string
		permissions map[string]string
	}

	var items = map[models.PermissionScope]Acl{
		&models.User{}: {
			title: "Manage User",
			permissions: map[string]string{
				"view":   "View Users",
				"update": "Update Users",
				"delete": "Delete Users",
			},
		},
	}

	var admin models.User
	initializers.DB.First(&admin, "email = ?", models.AdminEmail)

	for model, value := range items {
		role := models.Role{Name: value.title}

		scope := model.GetScope()

		initializers.DB.FirstOrCreate(&role)

		for name, desc := range value.permissions {
			// generate name
			permissionName := scope + "." + name
			permission := models.Permission{Name: permissionName, Description: desc}

			// create permission
			initializers.DB.Where(models.Permission{Name: permissionName}).FirstOrCreate(&permission)

			// append
			initializers.DB.Model(&role).Association("Permissions").Append(&permission)
		}
		initializers.DB.Model(&admin).Association("Roles").Append(&role)

	}

}
