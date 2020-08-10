package ir

import (
	"github.com/shu22203/internet_ranking/challenge"
	"github.com/shu22203/internet_ranking/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockSectionChallenges ChallengeResults

func (mas MockSectionChallenges) AggregateSubmissions() ChallengeResults {
	return ChallengeResults(mas)
}

func TestAggregateChallengeResults(t *testing.T) {
	t.Run("同一userの課題結果は一つの個人部門結果に統合される", func(t *testing.T) {
		user1Id := user.NewUserId()
		user2Id := user.NewUserId()
		challengeId1 := challenge.NewChallengeId()
		challengeId2 := challenge.NewChallengeId()

		section := Section{
			SectionChallenges: MockSectionChallenges(
				ChallengeResults{
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
