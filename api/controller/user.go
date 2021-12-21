package controller

import (
	"otter-v2/api/dao"
	"otter-v2/api/http/response"
	"otter-v2/api/middleware"
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
