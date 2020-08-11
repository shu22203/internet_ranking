package ir

import (
	"github.com/shu22203/internet_ranking/challenge"
)

type SectionChallenge interface {
	AwardPointsToSubmitters() challenge.ChallengeResult
}

type SectionChallenges map[challenge.ChallengeId]SectionChallenge

func (sc *SectionChallenges) AggregateChallengeResults() SectionResult {
	result := make(SectionResult)
	for cid, v := range *sc {
		for uid, point := range v.AwardPointsToSubmitters() {
			v, ok := result[uid]

			if ok {
				v.add(cid, point)
			} else {
				result[uid] = newIndividualSectionResult(cid, point)
			}
		}
	}

	return result
}
