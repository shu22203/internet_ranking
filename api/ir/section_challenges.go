package ir

import (
	"github.com/shu22203/internet_ranking/challenge"
)

type ChallengeResults map[challenge.ChallengeId]challenge.ChallengeResult

type SectionChallenges interface {
	AggregateSubmissions() ChallengeResults
}

type SectionChallengesImpl map[challenge.ChallengeId]challenge.Challenge

func (sc *SectionChallengesImpl) AggregateSubmissions() ChallengeResults {
	ret := make(ChallengeResults)
	for k, v := range *sc {
		ret[k] = v.AwardPointsToSubmitters()
	}

	return ret
}
