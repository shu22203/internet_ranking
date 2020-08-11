package challenge

import (
	"github.com/shu22203/internet_ranking/user"
)

type SubmissionType struct {
	coefficient float64
}

type Submission struct {
	userId         user.UserId
	submissionType SubmissionType
	score          int
}

func (s *Submission) ChallengeScore() float64 {
	return float64(s.score) * s.submissionType.coefficient
}
