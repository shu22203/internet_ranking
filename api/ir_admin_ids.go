package main

type IrAdminIds struct {
	ownerId  UserId
	adminIds []UserId
}

func NewIrAdminIds(owner *User) IrAdminIds {
	return IrAdminIds{
		ownerId:  owner.Id,
		adminIds: []UserId{},
	}
}

func (ira *IrAdminIds) AdminIds() []UserId {
	return append([]UserId{ira.ownerId}, ira.adminIds...)
}

func (ira *IrAdminIds) OwnerId() UserId {
	return ira.ownerId
}
