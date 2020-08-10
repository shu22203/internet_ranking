package main

import (
	"github.com/google/uuid"
)

type ChallengeId uuid.UUID

func NewChallengeId() ChallengeId {
	return ChallengeId(uuid.Must(uuid.NewRandom()))
}
