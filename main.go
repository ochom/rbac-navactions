package main

import (
	"fmt"

	"example.com/ochom/hello/models"
)

func AllMenu() []models.Menu {
	return []models.Menu{
		{Code: "001", Title: "Home", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: true, Requires: models.PermissionViewConsumer},
		{Code: "001", Title: "Consumers", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: true, Requires: models.PermissionViewConsumer},
		{Code: "001", Title: "Update Consumer", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: false, Requires: models.PermissionUpdateConsumer},
		{Code: "001", Title: "Delete Consumer", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: false, Requires: models.PermissionDeleteConsumer},
		{Code: "002", Title: "Agents", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: false, Requires: models.PermissionActivateAgent},
		{Code: "002", Title: "Agent Registration", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: true, Requires: models.PermissionActivateAgent},
		{Code: "002", Title: "Agent Identification", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: true, Requires: models.PermissionActivateAgent},
		{Code: "003", Title: "Patients", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: false, Requires: models.PermissionActivateAgent},
		{Code: "003", Title: "Identification", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: false, Requires: models.PermissionActivateAgent},
		{Code: "003", Title: "Registration", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: false, Requires: models.PermissionActivateAgent},
		{Code: "004", Title: "Profile", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: false, Requires: models.PermissionActivateAgent},
	}
}

func GetNavigationActions(u models.User) []models.Menu {
	allNavActions := AllMenu()
	userActions := []models.Menu{}
	for _, v := range allNavActions {
		u.HasPermission(v.Requires)
		userActions = append(userActions, v)
	}
	return userActions
}

func GroupNested(actions []models.Menu) []models.Menu {
	grouped := []models.Menu{}
	for _, v := range actions {
		if v.IsParent {
			for _, x := range actions {
				if !x.IsParent && x.Code == v.Code {
					v.Nested = append(v.Nested, x)
				}
			}
			grouped = append(grouped, v)
		}
	}
	return grouped
}

func GroupPriority(groupNested []models.Menu) (primary, secondary []models.Menu) {
	primary = []models.Menu{}
	secondary = []models.Menu{}

	added := make(map[string]models.Menu)

	if len(groupNested) < 5 {
		primary = groupNested
	} else {
		for {
			if len(primary) == 4 {
				break
			}
			// add all the high priority first
			for _, v := range groupNested {
				if v.Priority == models.MenuPriorityHigh {
					_, ok := added[v.Code]
					if !ok {
						primary = append(primary, v)
						added[v.Code] = v
					}
				}
			}
			// add every other item
			for _, v := range groupNested {
				_, ok := added[v.Code]
				if !ok {
					primary = append(primary, v)
					added[v.Code] = v
				}
			}
		}
		// add all remaining items to secondary
		for _, v := range groupNested {
			_, ok := added[v.Code]
			if !ok {
				secondary = append(secondary, v)
				added[v.Code] = v
			}
		}
	}

	for _, v := range groupNested {
		if len(primary) < 4 {
			if v.Priority == models.MenuPriorityHigh {
				primary = append(primary, v)
				added[v.Code] = v
			}
		}
	}
	return primary, secondary
}

func main() {
	role := models.RoleTypeAgent
	user := models.User{
		Roles:        []models.Role{role},
		Permissions:  role.Permissions(),
		FavoriteMenu: []string{"Home"},
	}

	userNavActions := GetNavigationActions(user)
	groupNested := GroupNested(userNavActions)
	primary, secondary := GroupPriority(groupNested)

	fmt.Println(primary, secondary)
}
