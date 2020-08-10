package challenge

import (
	"github.com/shu22203/internet_ranking/user"
)

type SubmissionType struct {
	Coefficient float64
}

type Submission struct {
	UserId         user.UserId
	SubmissionType SubmissionType
	Score          int
}

func (s *Submission) ChallengeScore() float64 {
	return float64(s.Score) * s.SubmissionType.Coefficient
}
