package delivery

import (
	"otter-v2/api/middleware"
	"otter-v2/api/response"
	"otter-v2/app/usecase"
)

var User = userDelivery{}

type userDelivery struct{}

func (d userDelivery) GetUserList(webInput middleware.WebInput) {
	userList, err := usecase.User.GetUserList()

	if err != nil {
		response.ServerError(webInput.Context.Ctx, err.Error(), err)
		return
	}
	response.Success(webInput.Context.Ctx, userList)
}
