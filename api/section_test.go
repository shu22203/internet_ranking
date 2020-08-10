package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockSectionChallenges ChallengeResult

func (mas MockSectionChallenges) AggregateSubmissions() ChallengeResult {
	return ChallengeResult(mas)
}

func TestAggregateChallengeResults(t *testing.T) {
	t.Run("同一userの課題結果は一つの個人部門結果に統合される", func(t *testing.T) {
		user1Id := NewUserId()
		user2Id := NewUserId()
		challengeId1 := NewChallengeId()
		challengeId2 := NewChallengeId()

		section := Section{
			SectionChallenges: MockSectionChallenges(
				ChallengeResult{
					challengeId1: {
						user1Id: 12,
						user2Id: 15,
					},
					challengeId2: {
						user1Id: 15,
					},
				},
			),
		}

		expect := SectionResult{
			user1Id: {
				challengeId1: 12,
				challengeId2: 15,
			},
			user2Id: {
				challengeId1: 15,
			},
		}

		assert.Equal(t, expect, section.AggregateChallengeResults())
	})
}
