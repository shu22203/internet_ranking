package main

import (
	"github.com/shu22203/internet_ranking/challenge"
	"github.com/shu22203/internet_ranking/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAggregateSubmissions(t *testing.T) {
	userId1 := user.NewUserId()
	userId2 := user.NewUserId()
	userId3 := user.NewUserId()
	challengeId1 := challenge.NewChallengeId()
	challengeId2 := challenge.NewChallengeId()
	challengeId3 := challenge.NewChallengeId()

	scs := SectionChallengesImpl{
		challengeId1: {
			MaxPoint: 10,
			MinPoint: 0,
			Submissions: []challenge.Submission{
				{
					UserId:         userId1,
					SubmissionType: challenge.SubmissionType{Coefficient: 0.9},
					Score:          1000000,
				},
				{
					UserId:         userId2,
					SubmissionType: challenge.SubmissionType{Coefficient: 1.0},
					Score:          1000000,
				},
				{
					UserId:         userId3,
					SubmissionType: challenge.SubmissionType{Coefficient: 1.0},
					Score:          850000,
				},
			},
		},
		challengeId2: {
			MaxPoint: 10,
			MinPoint: 0,
			Submissions: []challenge.Submission{
				{
					UserId:         userId1,
					SubmissionType: challenge.SubmissionType{Coefficient: 1.0},
					Score:          1000000,
				},
				{
					UserId:         userId3,
					SubmissionType: challenge.SubmissionType{Coefficient: 0.9},
					Score:          850000,
				},
			},
		},
		challengeId3: {
			MaxPoint: 5,
			MinPoint: 5,
			Submissions: []challenge.Submission{
				{
					UserId:         userId1,
					SubmissionType: challenge.SubmissionType{Coefficient: 1.0},
					Score:          1,
				},
				{
					UserId:         userId2,
					SubmissionType: challenge.SubmissionType{Coefficient: 1.0},
					Score:          1,
				},
			},
		},
	}

	expect := ChallengeResults{
		challengeId1: {
			userId1: 9,
			userId2: 10,
			userId3: 8,
		},
		challengeId2: {
			userId1: 10,
			userId3: 9,
		},
		challengeId3: {
			userId1: 5,
			userId2: 5,
		},
	}

	assert.Equal(t, expect, scs.AggregateSubmissions())
}
