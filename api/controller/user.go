package event

import (
	"otter-calendar-ws/api/dao"
	"otter-calendar-ws/service"
)

var User = userController{}

type userController struct{}

func (controller userController) GetUserList(webInput interceptor.WebInput) {
	userList := dao.User.GetEventList()
	service.Response.OK(webInput.userList)
}
