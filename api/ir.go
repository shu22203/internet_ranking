package main

type InternetRanking struct {
	IrAdminIds
	IrSections
}

func HoldInternetRanking(user *User) *InternetRanking {
	return &InternetRanking{
		IrAdminIds: NewIrAdminIds(user),
	}
}
