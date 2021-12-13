package controller

import (
	"otter-calendar-ws/api/dao"
	"otter-calendar-ws/api/http/response"
	"otter-calendar-ws/api/middleware"
)

var User = userController{}

type userController struct{}

func (controller userController) GetUserList(webInput middleware.WebInput) {
	userList := dao.User.GetEventList()
	response.Success(webInput.Context.Ctx, userList)
}
