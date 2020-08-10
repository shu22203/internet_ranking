package main

type SubmissionType struct {
	coefficient float64
}

type Submission struct {
	userId         UserId
	submissionType SubmissionType
	score          int
}

func (s *Submission) ChallengeScore() float64 {
	return float64(s.score) * s.submissionType.coefficient
}
