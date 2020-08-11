package ir

import (
	"github.com/shu22203/internet_ranking/challenge"
	"github.com/shu22203/internet_ranking/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockSectionChallenge challenge.ChallengeResult

func (msc MockSectionChallenge) AwardPointsToSubmitters() challenge.ChallengeResult {
	return challenge.ChallengeResult(msc)
}

func TestAggregateChallengeResults(t *testing.T) {
	t.Run("同一userの課題結果は一つの個人部門結果に統合される", func(t *testing.T) {
		user1Id := user.NewUserId()
		user2Id := user.NewUserId()
		challengeId1 := challenge.NewChallengeId()
		challengeId2 := challenge.NewChallengeId()

		section := Section{
			challenges: map[challenge.ChallengeId]SectionChallenge{
				challengeId1: MockSectionChallenge{
					user1Id: 12,
					user2Id: 15,
				},
				challengeId2: MockSectionChallenge{
					user1Id: 15,
				},
			},
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
