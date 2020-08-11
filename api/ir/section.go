package ir

import (
	"github.com/shu22203/internet_ranking/challenge"
	"time"
)

type Section struct {
	challenges map[challenge.ChallengeId]SectionChallenge
	startAt    time.Time
	endAt      time.Time
}

func (s *Section) AggregateChallengeResults() SectionResult {
	crs := make(map[challenge.ChallengeId]challenge.ChallengeResult)
	for k, v := range s.challenges {
		crs[k] = v.AwardPointsToSubmitters()
	}

	result := make(SectionResult)
	for cid, cr := range crs {
		for uid, point := range cr {
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
