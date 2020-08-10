package main

type SectionChallenges interface {
	AbleToAggregateSubmissions
}

type AbleToAggregateSubmissions interface {
	AggregateSubmissions() ChallengeResult
}

type ChallengeResult map[ChallengeId]map[UserId]int
