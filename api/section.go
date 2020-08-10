package main

import (
	"time"
)

type Section struct {
	startAt time.Time
	endAt   time.Time
	AbleToAggregateSubmissions
}

func (s *Section) AggregateChallengeResults() SectionResult {
	result := make(SectionResult)
	for cid, cr := range s.AggregateSubmissions() {
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
