package main

import (
	"github.com/google/uuid"
	"sort"
)

type ChallengeId uuid.UUID

func NewChallengeId() ChallengeId {
	return ChallengeId(uuid.Must(uuid.NewRandom()))
}

type Challenge struct {
	maxPoint    int
	minPoint    int
	submissions []Submission
}

func (c *Challenge) AwardPointsToSubmitters() map[UserId]int {
	ss := []Submission{}
	for _, as := range c.AdoptedSubmissions() {
		ss = append(ss, as)
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].ChallengeScore() > ss[j].ChallengeScore()
	})

	ret := make(map[UserId]int)
	for i, s := range ss {
		ret[s.userId] = c.maxPoint - i

		if ret[s.userId] < c.minPoint {
			ret[s.userId] = c.minPoint
		}

		if i != 0 && s.ChallengeScore() == ss[i-1].ChallengeScore() {
			ret[s.userId] = ret[ss[i-1].userId]
		}
	}

	return ret
}

func (c *Challenge) AdoptedSubmissions() map[UserId]Submission {
	ret := make(map[UserId]Submission)
	for k, v := range c.eachUserSubmissions() {
		maxCs := v[0].ChallengeScore()
		ret[k] = v[0]
		for _, sub := range v {
			if sub.ChallengeScore() > maxCs {
				maxCs = sub.ChallengeScore()
				ret[k] = sub
			}
		}

	}

	return ret
}

func (c *Challenge) eachUserSubmissions() map[UserId][]Submission {
	ret := make(map[UserId][]Submission)
	for _, s := range c.submissions {
		v, ok := ret[s.userId]

		if ok {
			ret[s.userId] = append(v, s)
		} else {
			ret[s.userId] = []Submission{s}
		}
	}

	return ret
}
