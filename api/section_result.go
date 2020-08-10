package main

type SectionResult map[UserId]*SectionIndividualResult

type SectionIndividualResult map[ChallengeId]int

func newIndividualSectionResult(cid ChallengeId, point int) *SectionIndividualResult {
	return &SectionIndividualResult{
		cid: point,
	}
}

func (sir *SectionIndividualResult) add(cid ChallengeId, point int) {
	(*sir)[cid] = point
}

func (sir *SectionIndividualResult) TotalPoint() int {
	point := 0
	for _, p := range *sir {
		point += p
	}

	return point
}
