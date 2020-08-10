package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockAggregateSubmissions ChallengeResult

func (mas MockAggregateSubmissions) AggregateSubmissions() ChallengeResult {
	return ChallengeResult(mas)
}

func TestAggregateChallengeResults(t *testing.T) {
	t.Run("同一userの課題結果は一つの個人部門結果に統合される", func(t *testing.T) {
		user1Id := NewUserId()
		user2Id := NewUserId()
		challengeId1 := NewChallengeId()
		challengeId2 := NewChallengeId()

		expect := SectionResult{
			user1Id: {
				challengeId1: 12,
				challengeId2: 15,
			},
			user2Id: {
				challengeId1: 15,
			},
		}

		section := Section{
			AbleToAggregateSubmissions: MockAggregateSubmissions(
				ChallengeResult{
					challengeId1: map[UserId]int{
						user1Id: 12,
						user2Id: 15,
					},
					challengeId2: map[UserId]int{
						user1Id: 15,
					},
				},
			),
		}

		actual := section.AggregateChallengeResults()

		assert.Equal(t, expect, actual)
	})
}
