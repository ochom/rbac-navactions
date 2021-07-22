package models

var PermissionViewConsumer Permission = Permission{Resource: "consumer", Action: "view"}
var PermissionEditConsumer Permission = Permission{Resource: "consumer", Action: "edit"}
var PermissionUpdateConsumer Permission = Permission{Resource: "consumer", Action: "update"}
var PermissionDeleteConsumer Permission = Permission{Resource: "consumer", Action: "delete"}
var PermissionActivateAgent Permission = Permission{Resource: "agent", Action: "activate"}

var AgentPerms []Permission = []Permission{
	PermissionViewConsumer,
	PermissionEditConsumer,
}
