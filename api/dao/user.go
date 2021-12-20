package dao

import (
	"otter-calendar-ws/api/model/user"
	"otter-calendar-ws/jobqueue"
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
