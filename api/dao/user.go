package dao

import (
	"otter-calendar-ws/entity"
	"otter-calendar-ws/model"
	"otter-calendar-ws/service"
)

var User = userDao{
	Table: service.Model.GetTable(&model.User{}),
	PK:    service.Model.GetPK(&model.User{}),
	Col:   service.Model.GetColumn(&model.User{}).(model.User),
}

type userDao struct {
	Table string
	PK    string
	Col   model.User
}

func (dao userDao) GetEventList() []entity.User {
	return []entity.User{}
}
