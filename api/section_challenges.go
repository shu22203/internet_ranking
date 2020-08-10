package main

type SectionChallenges interface {
	AggregateSubmissions() ChallengeResult
}

type ChallengeResult map[ChallengeId]map[UserId]int
