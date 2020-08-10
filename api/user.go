package main

import (
	"github.com/google/uuid"
)

type UserId uuid.UUID

func NewUserId() UserId {
	return UserId(uuid.Must(uuid.NewRandom()))
}

type User struct {
	Id UserId
}

func NewUser() *User {
	return &User{
		Id: NewUserId(),
	}
}
