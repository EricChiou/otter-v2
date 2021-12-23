package router

import (
	"otter-v2/app/delivery"
)

func initUserAPI() {
	groupName := "/user"

	// user list
	get(groupName+"/list", false, delivery.User.GetUserList)
}
