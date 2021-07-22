package main

import (
	"encoding/json"
	"net/http"

	"example.com/ochom/hello/models"
)

func GetNavigationActions(u models.User) []models.Menu {
	userActions := []models.Menu{}
	for _, action := range models.AllNavActions {
		if u.HasPermission(action.Requires) {
			if u.HasFavoriteMenu(action.Title) {
				action.Favourite = true
			}
			userActions = append(userActions, action)
		}
	}
	return userActions
}

func GroupNested(actions []models.Menu) []models.Menu {
	grouped := []models.Menu{}
	for _, parent := range actions {
		if parent.IsParent {
			for _, child := range actions {
				if !child.IsParent && child.Code == parent.Code {
					parent.Nested = append(parent.Nested, child)
				}
			}
			grouped = append(grouped, parent)
		}
	}
	return grouped
}

func GroupPriority(actions []models.Menu) (primary, secondary []models.Menu) {
	primary = []models.Menu{
		{
			Code:       "001",
			Title:      "Home",
			Icon:       models.Link{URL: "test.url.com/png"},
			OnTapRoute: "",
			IsParent:   true,
			Requires:   models.PermissionViewConsumer,
			Primary:    true,
		},
	}

	secondary = []models.Menu{}

	added := make(map[string]models.Menu)

	//pb is number of navactions that can possibly be on bottom navigation
	pb := 0
	for _, v := range actions {
		if len(v.Nested) == 0 {
			pb += 1
		}
	}

	// add all the possible bottom action to primary is they are less or equal to 4
	if pb <= 3 {
		for _, action := range actions {
			if len(action.Nested) == 0 {
				primary = append(primary, action)
				added[action.Code] = action
			}
		}
	} else {
		for {
			if len(primary) == 4 {
				break
			}
			// // add all the high priority first
			for _, action := range actions {
				if action.Primary {

					_, exist := added[action.Code]
					if !exist && len(action.Nested) == 0 {
						primary = append(primary, action)
						added[action.Code] = action
					}

					if len(primary) == 4 {
						break
					}
				}
			}
		}
	}

	// add all remaining items to secondary
	for _, action := range actions {
		_, exists := added[action.Code]
		if !exists {
			secondary = append(secondary, action)
			added[action.Code] = action
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
