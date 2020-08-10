package main

type SectionChallenges interface {
	AggregateSubmissions() ChallengeResults
}

type ChallengeResults map[ChallengeId]map[UserId]int
