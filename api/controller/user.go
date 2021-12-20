package controller

import (
	"otter-calendar-ws/api/dao"
	"otter-calendar-ws/api/http/response"
	"otter-calendar-ws/api/middleware"
)

var User = userController{}

type userController struct{}

func (controller userController) GetUserList(webInput middleware.WebInput) {
	userList, err := dao.User.GetEventList()

	if err != nil {
		response.Error(webInput.Context.Ctx, response.ServerError, err.Error(), err)
		return
	}
	response.Success(webInput.Context.Ctx, userList)
}
