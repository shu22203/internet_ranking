package main

type ChallengeResults map[ChallengeId]map[UserId]int

type SectionChallenges interface {
	AggregateSubmissions() ChallengeResults
}

type SectionChallengesImpl map[ChallengeId]Challenge

func (sc *SectionChallengesImpl) AggregateSubmissions() ChallengeResults {
	ret := make(ChallengeResults)
	for k, v := range *sc {
		ret[k] = v.AwardPointsToSubmitters()
	}

	return ret
}
