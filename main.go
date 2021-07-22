package main

import (
	"encoding/json"
	"net/http"

	"example.com/ochom/hello/models"
)

func AllMenu() []models.Menu {
	return []models.Menu{
		{Code: "001", Title: "Home", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: true, Requires: models.PermissionViewConsumer, Priority: models.MenuPriorityHigh},
		{Code: "002", Title: "Consumers", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: true, Requires: models.PermissionViewConsumer, Priority: models.MenuPriorityHigh},
		{Code: "002", Title: "Update Consumer", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: false, Requires: models.PermissionUpdateConsumer},
		{Code: "002", Title: "Delete Consumer", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: false, Requires: models.PermissionDeleteConsumer},
		{Code: "003", Title: "Agents", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: true, Requires: models.PermissionActivateAgent, Priority: models.MenuPriorityLow},
		{Code: "003", Title: "Agent Registration", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: false, Requires: models.PermissionActivateAgent},
		{Code: "003", Title: "Agent Identification", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: false, Requires: models.PermissionActivateAgent},
		{Code: "004", Title: "Patients", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: true, Requires: models.PermissionActivateAgent, Priority: models.MenuPriorityHigh},
		{Code: "004", Title: "Identification", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: false, Requires: models.PermissionActivateAgent},
		{Code: "004", Title: "Registration", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: false, Requires: models.PermissionActivateAgent},
		{Code: "005", Title: "Profile", Icon: models.Link{URL: "test.url.com/png"}, OnTapRoute: "", IsParent: true, Requires: models.PermissionActivateAgent, Priority: models.MenuPriorityLow},
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
		for len(primary) < 4 {
			// // add all the high priority first
			for _, v := range groupNested {
				if v.Priority == models.MenuPriorityHigh {
					_, ok := added[v.Code]
					if !ok {
						primary = append(primary, v)
						added[v.Code] = v
					}

					if len(primary) == 4 {
						break
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

				if len(primary) == 4 {
					break
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
	return primary, secondary
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	role := models.RoleTypeAgent
	user := models.User{
		Roles:        []models.Role{role},
		Permissions:  role.Permissions(),
		FavoriteMenu: []string{"Home"},
	}

	userNavActions := GetNavigationActions(user)
	groupNested := GroupNested(userNavActions)
	primary, secondary := GroupPriority(groupNested)
	data := struct {
		Primary   []models.Menu `json:"primary"`
		Secondary []models.Menu `json:"secondary"`
	}{
		Primary:   primary,
		Secondary: secondary,
	}
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8081", nil)
}
