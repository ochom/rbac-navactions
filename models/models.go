package models

type Link struct {
	URL string `json:"url,omitempty"`
}

type Menu struct {
	Code       string        `json:"code,omitempty"`
	Title      string        `json:"title,omitempty"`
	OnTapRoute string        `json:"onTapRoute,omitempty"`
	Icon       Link          `json:"icon,omitempty"`
	Favourite  bool          `json:"favourite,omitempty"`
	IsParent   bool          `json:"isParent,omitempty"`
	Nested     []interface{} `json:"nested,omitempty"`
	Requires   Permission    `json:"requires,omitempty"`
	Primary    bool          `json:"primary"`
}

type Role string

const (
	RoleTypeAgent         Role = "Agent"
	RoleTypeEmployee      Role = "Employee"
	RoleTypeAdministrator Role = "Administrator"
)

type Permission struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

type User struct {
	Roles        []Role   `json:"roles,omitempty"`
	FavoriteMenu []string `json:"favoriteMenu,omitempty"`
}
