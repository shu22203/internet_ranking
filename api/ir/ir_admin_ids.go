package ir

import (
	"github.com/shu22203/internet_ranking/user"
)

type IrAdminIds struct {
	ownerId  user.UserId
	adminIds []user.UserId
}

func NewIrAdminIds(owner *user.User) IrAdminIds {
	return IrAdminIds{
		ownerId:  owner.Id,
		adminIds: []user.UserId{},
	}
}

func (ira *IrAdminIds) AdminIds() []user.UserId {
	return append([]user.UserId{ira.ownerId}, ira.adminIds...)
}

func (ira *IrAdminIds) OwnerId() user.UserId {
	return ira.ownerId
}
