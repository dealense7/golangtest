package seeds

import (
	models2 "github.com/dealense7/documentSignatures/app/models"
	"github.com/dealense7/documentSignatures/initializers"
)

func AclSeed() {

	type Acl struct {
		title       string
		permissions map[string]string
	}

	var items = map[models2.PermissionScope]Acl{
		&models2.User{}: {
			title: "Manage User",
			permissions: map[string]string{
				"view":   "View Users",
				"update": "Update Users",
				"delete": "Delete Users",
			},
		},
	}

	var admin models2.User
	initializers.DB.First(&admin, "email = ?", models2.AdminEmail)

	for model, value := range items {
		role := models2.Role{Name: value.title}

		scope := model.GetScope()

		initializers.DB.FirstOrCreate(&role)

		for name, desc := range value.permissions {
			// generate name
			permissionName := scope + "." + name
			permission := models2.Permission{Name: permissionName, Description: desc}

			// create permission
			initializers.DB.Where(models2.Permission{Name: permissionName}).FirstOrCreate(&permission)

			// append
			initializers.DB.Model(&role).Association("Permissions").Append(&permission)
		}
		initializers.DB.Model(&admin).Association("Roles").Append(&role)

	}

}
