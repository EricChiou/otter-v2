package dao

import (
	"otter-v2/api/model/user"
	"otter-v2/jobqueue"
)

var User = userDao{}

type userDao struct{}

func (dao userDao) GetEventList() ([]user.Entity, error) {
	var userList []user.Entity
	err := jobqueue.User.NewUserListJob(func() interface{} {
		userList = []user.Entity{}
		return nil
	})
	return userList, err
}
