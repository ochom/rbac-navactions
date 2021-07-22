package models

func (u User) HasRole(role Role) bool {
	for _, v := range u.Roles {
		if v == role {
			return true
		}
	}
	return false
}

func (u User) HasPermission(perm Permission) bool {
	perms := []Permission{}

	// get all permissions
	for _, r := range u.Roles {
		perms = append(perms, r.Permissions()...)
	}

	// check if required permission exists
	for _, v := range perms {
		if v == perm {
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
