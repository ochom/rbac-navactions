package models

var PermissionViewConsumer Permission = Permission{Resource: "consumer", Action: "view"}
var PermissionEditConsumer Permission = Permission{Resource: "consumer", Action: "edit"}
var PermissionUpdateConsumer Permission = Permission{Resource: "consumer", Action: "update"}
var PermissionDeleteConsumer Permission = Permission{Resource: "consumer", Action: "delete"}
var PermissionActivateAgent Permission = Permission{Resource: "agent", Action: "activate"}

var AgentPerms []Permission = []Permission{
	PermissionViewConsumer,
	PermissionEditConsumer,
	PermissionUpdateConsumer,
}

var AllNavActions []Menu = []Menu{

	{
		Code:       "002",
		Title:      "Consumers",
		Icon:       Link{URL: "test.url.com/png"},
		OnTapRoute: "",
		IsParent:   true,
		Requires:   PermissionViewConsumer,
		Priority:   MenuPriorityHigh,
	},
	{
		Code:       "002",
		Title:      "Update Consumer",
		Icon:       Link{URL: "test.url.com/png"},
		OnTapRoute: "",
		IsParent:   false,
		Requires:   PermissionUpdateConsumer,
	},
	{
		Code:       "002",
		Title:      "Delete Consumer",
		Icon:       Link{URL: "test.url.com/png"},
		OnTapRoute: "",
		IsParent:   false,
		Requires:   PermissionDeleteConsumer,
	},
	{
		Code:       "003",
		Title:      "Agents",
		Icon:       Link{URL: "test.url.com/png"},
		OnTapRoute: "",
		IsParent:   true,
		Requires:   PermissionActivateAgent,
		Priority:   MenuPriorityLow,
	},
	{
		Code:       "003",
		Title:      "Agent Registration",
		Icon:       Link{URL: "test.url.com/png"},
		OnTapRoute: "",
		IsParent:   false,
		Requires:   PermissionActivateAgent,
	},
	{
		Code:       "003",
		Title:      "Agent Identification",
		Icon:       Link{URL: "test.url.com/png"},
		OnTapRoute: "",
		IsParent:   false,
		Requires:   PermissionActivateAgent,
	},
	{
		Code:       "004",
		Title:      "Patients",
		Icon:       Link{URL: "test.url.com/png"},
		OnTapRoute: "",
		IsParent:   true,
		Requires:   PermissionActivateAgent,
		Priority:   MenuPriorityHigh,
	},
	{
		Code:       "004",
		Title:      "Identification",
		Icon:       Link{URL: "test.url.com/png"},
		OnTapRoute: "",
		IsParent:   false,
		Requires:   PermissionActivateAgent,
	},
	{
		Code:       "004",
		Title:      "Registration",
		Icon:       Link{URL: "test.url.com/png"},
		OnTapRoute: "",
		IsParent:   false,
		Requires:   PermissionActivateAgent,
	},
	{
		Code:       "005",
		Title:      "Profile",
		Icon:       Link{URL: "test.url.com/png"},
		OnTapRoute: "",
		IsParent:   true,
		Requires:   PermissionActivateAgent,
		Priority:   MenuPriorityLow,
	},
}
