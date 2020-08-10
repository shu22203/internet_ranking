package main

import (
	"github.com/shu22203/internet_ranking/challenge"
	"github.com/shu22203/internet_ranking/user"
)

type SectionResult map[user.UserId]*SectionIndividualResult

type SectionIndividualResult map[challenge.ChallengeId]int

func newIndividualSectionResult(cid challenge.ChallengeId, point int) *SectionIndividualResult {
	return &SectionIndividualResult{
		cid: point,
	}
}

func (sir *SectionIndividualResult) add(cid challenge.ChallengeId, point int) {
	(*sir)[cid] = point
}

func (sir *SectionIndividualResult) TotalPoint() int {
	point := 0
	for _, p := range *sir {
		point += p
	}

	return point
}
