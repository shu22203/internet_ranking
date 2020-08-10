package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAggregateSubmissions(t *testing.T) {
	userId1 := NewUserId()
	userId2 := NewUserId()
	userId3 := NewUserId()
	challengeId1 := NewChallengeId()
	challengeId2 := NewChallengeId()
	challengeId3 := NewChallengeId()

	scs := SectionChallengesImpl{
		challengeId1: {
			maxPoint: 10,
			minPoint: 0,
			submissions: []Submission{
				{
					userId: userId1,
					submissionType: SubmissionType{coefficient: 0.9},
					score: 1000000,
				},
				{
					userId: userId2,
					submissionType: SubmissionType{coefficient: 1.0},
					score: 1000000,
				},
				{
					userId: userId3,
					submissionType: SubmissionType{coefficient: 1.0},
					score: 850000,
				},
			},
		},
		challengeId2: {
			maxPoint: 10,
			minPoint: 0,
			submissions: []Submission{
				{
					userId: userId1,
					submissionType: SubmissionType{coefficient: 1.0},
					score: 1000000,
				},
				{
					userId: userId3,
					submissionType: SubmissionType{coefficient: 0.9},
					score: 850000,
				},
			},
		},
		challengeId3: {
			maxPoint: 5,
			minPoint: 5,
			submissions: []Submission{
				{
					userId: userId1,
					submissionType: SubmissionType{coefficient: 1.0},
					score: 1,
				},
				{
					userId: userId2,
					submissionType: SubmissionType{coefficient: 1.0},
					score: 1,
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
