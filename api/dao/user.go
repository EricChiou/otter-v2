package dao

import "otter-calendar-ws/api/model/user"

var User = userDao{}

type userDao struct{}

func (dao userDao) GetEventList() []user.Entity {
	return []user.Entity{}
}
