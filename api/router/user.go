package router

import "otter-calendar-ws/api/controller"

func initUserAPI() {
	groupName := "/user"
	controller := controller.User

	// user list
	get(groupName+"/list", false, controller.GetUserList)
}
