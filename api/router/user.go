package router

import "otter-v2/api/controller"

func initUserAPI() {
	groupName := "/user"
	controller := controller.User

	// user list
	get(groupName+"/list", false, controller.GetUserList)
}
