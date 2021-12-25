package repository

import (
	"otter-v2/app/model/user"
	"otter-v2/service/jobqueue"
)

var User = userRepository{}

type userRepository struct{}

func (r userRepository) GetUserList() ([]user.Entity, error) {
	var userList []user.Entity
	err := jobqueue.User.NewUserListJob(func() interface{} {
		userList = []user.Entity{}
		return nil
	})

	return userList, err
}
