package models

func (u User) HasRole(role Role) bool {
	for _, v := range u.Roles {
		if v == role {
			return true
		}
	}
	return false
}

func (u User) GetPermissions() []Permission {
	perms := []Permission{}

	// get all permissions
	for _, role := range u.Roles {
		perms = append(perms, role.Permissions()...)
	}

	// remove duplicated perms

	return perms
}

func (u User) HasPermission(perm Permission) bool {
	permissions := u.GetPermissions()

	// check if required permission exists
	for _, permission := range permissions {
		if permission == perm {
			return true
		}
	}
	return false
}

func (u User) HasFavoriteMenu(title string) bool {
	for _, v := range u.FavoriteMenu {
		if v == title {
			return true
		}
	}
	return false
}

// Permission checks
func (r Role) Permissions() []Permission {
	switch r {
	case RoleTypeAgent:
		return AgentPerms
	}
	return []Permission{}
}
