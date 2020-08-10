package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdpotedSubmissions(t *testing.T) {
	t.Run("同一userの提出は係数を掛けたスコアが最も高い提出を採用する", func(t *testing.T) {
		userId := NewUserId()
		challenge := Challenge{
			submissions: []Submission{
				{
					userId: userId,
					submissionType: SubmissionType{
						coefficient: 0.8,
					},
					score: 1000000,
				},
				{
					userId: userId,
					submissionType: SubmissionType{
						coefficient: 0.9,
					},
					score: 980000,
				},
				{
					userId: userId,
					submissionType: SubmissionType{
						coefficient: 1.0,
					},
					score: 950000,
				},
			},
		}

		expect := map[UserId]Submission{
			userId: {
				userId: userId,
				submissionType: SubmissionType{
					coefficient: 1.0,
				},
				score: 950000,
			},
		}

		actual := challenge.AdoptedSubmissions()

		assert.Equal(t, expect, actual)
	})
}
