package delivery

import (
	"otter-v2/app/usecase"
	"otter-v2/http/middleware"
	"otter-v2/http/response"
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
