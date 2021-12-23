package usecase

import (
	"otter-v2/app/model/user"
	"otter-v2/app/repository"
)

var User userUsecase

type userUsecase struct{}

func (u userUsecase) GetUserList() ([]user.Entity, error) {
	userList, err := repository.User.GetUserList()
	return userList, err
}
