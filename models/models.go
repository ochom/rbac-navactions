package models

type Link struct {
	URL string `json:"url,omitempty"`
}

type MenuPriority string

const (
	MenuPriorityHigh MenuPriority = "HIGH"
	MenuPriorityLow  MenuPriority = "LOW"
)

type Menu struct {
	Code       string        `json:"code,omitempty"`
	Title      string        `json:"title,omitempty"`
	OnTapRoute string        `json:"onTapRoute,omitempty"`
	Icon       Link          `json:"icon,omitempty"`
	Favourite  bool          `json:"favourite,omitempty"`
	IsParent   bool          `json:"isParent,omitempty"`
	Nested     []interface{} `json:"nested,omitempty"`
	Requires   Permission    `json:"requires,omitempty"`
	Priority   MenuPriority  `json:"priority"`

	// Primary marks
	Primary bool `json:"primary"`
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
	Permissions  []Permission `json:"permissions,omitempty"`
	Roles        []Role       `json:"roles,omitempty"`
	FavoriteMenu []string     `json:"favoriteMenu,omitempty"`
}
