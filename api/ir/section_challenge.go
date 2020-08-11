package ir

import (
	"github.com/shu22203/internet_ranking/challenge"
)

type SectionChallenge interface {
	AwardPointsToSubmitters() challenge.ChallengeResult
}
