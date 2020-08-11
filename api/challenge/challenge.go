package challenge

import (
	"github.com/google/uuid"
	"github.com/shu22203/internet_ranking/user"
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

type ChallengeResult map[user.UserId]int

func (c *Challenge) AwardPointsToSubmitters() ChallengeResult {
	ss := []Submission{}
	for _, as := range c.AdoptedSubmissions() {
		ss = append(ss, as)
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].ChallengeScore() > ss[j].ChallengeScore()
	})

	ret := make(ChallengeResult)
	for i, s := range ss {
		ret[s.UserId] = c.maxPoint - i

		if ret[s.UserId] < c.minPoint {
			ret[s.UserId] = c.minPoint
		}

		if i != 0 && s.ChallengeScore() == ss[i-1].ChallengeScore() {
			ret[s.UserId] = ret[ss[i-1].UserId]
		}
	}

	return ret
}

func (c *Challenge) AdoptedSubmissions() map[user.UserId]Submission {
	ret := make(map[user.UserId]Submission)
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

func (c *Challenge) eachUserSubmissions() map[user.UserId][]Submission {
	ret := make(map[user.UserId][]Submission)
	for _, s := range c.submissions {
		v, ok := ret[s.UserId]

		if ok {
			ret[s.UserId] = append(v, s)
		} else {
			ret[s.UserId] = []Submission{s}
		}
	}

	return ret
}
