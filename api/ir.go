package main

import (
	"github.com/shu22203/internet_ranking/user"
)

type InternetRanking struct {
	IrAdminIds
	IrSections
}

func HoldInternetRanking(user *user.User) *InternetRanking {
	return &InternetRanking{
		IrAdminIds: NewIrAdminIds(user),
	}
}
