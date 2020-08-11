package challenge

import (
	"github.com/stretchr/testify/assert"
	"github.com/shu22203/internet_ranking/user"
	"testing"
)

func TestAdpotedSubmissions(t *testing.T) {
	t.Run("同一userの提出は課題スコアが最も高い提出を採用する", func(t *testing.T) {
		userId := user.NewUserId()
		challenge := Challenge{
			submissions: []Submission{
				{
					userId:         userId,
					submissionType: SubmissionType{coefficient: 0.8},
					score:          1000000,
				},
				{
					userId:         userId,
					submissionType: SubmissionType{coefficient: 0.9},
					score:          980000,
				},
				{
					userId:         userId,
					submissionType: SubmissionType{coefficient: 1.0},
					score:          950000,
				},
			},
		}

		expect := map[user.UserId]Submission{
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

func TestAwardPointsToSubmitters(t *testing.T) {
	t.Run("課題スコアの順位に応じたポイントを与える", func(t *testing.T) {
		userId1 := user.NewUserId()
		userId2 := user.NewUserId()
		userId3 := user.NewUserId()
		userId4 := user.NewUserId()
		userId5 := user.NewUserId()
		userId6 := user.NewUserId()
		userId7 := user.NewUserId()

		challenge := Challenge{
			maxPoint: 5,
			minPoint: 0,
			submissions: []Submission{
				{
					userId:         userId1,
					submissionType: SubmissionType{coefficient: 1.0},
					score:          500,
				},
				{
					userId:         userId2,
					submissionType: SubmissionType{coefficient: 1.0},
					score:          400,
				},
				{
					userId:         userId3,
					submissionType: SubmissionType{coefficient: 1.0},
					score:          300,
				},
				{
					userId:         userId4,
					submissionType: SubmissionType{coefficient: 1.0},
					score:          300,
				},
				{
					userId:         userId5,
					submissionType: SubmissionType{coefficient: 1.0},
					score:          200,
				},
				{
					userId:         userId6,
					submissionType: SubmissionType{coefficient: 1.0},
					score:          100,
				},
				{
					userId:         userId7,
					submissionType: SubmissionType{coefficient: 1.0},
					score:          0,
				},
			},
		}
		actual := challenge.AwardPointsToSubmitters()

		t.Run("1位の人が最大ポイントをもらえる", func(t *testing.T) {
			assert.Equal(t, 5, actual[userId1])
		})
		t.Run("順位が下がるごとに1点ずつもらえるポイントが減る", func(t *testing.T) {
			assert.Equal(t, 4, actual[userId2])
		})
		t.Run("3位同着なら両者とも3位相当の点数がもらえる", func(t *testing.T) {
			assert.Equal(t, 3, actual[userId3])
			assert.Equal(t, 3, actual[userId4])
		})
		t.Run("3位同着の次の順位の人は5位相当の点数がもらえる", func(t *testing.T) {
			assert.Equal(t, 1, actual[userId5])
		})
		t.Run("提出があれば最低ポイントは必ず取得できる", func(t *testing.T) {
			assert.Equal(t, 0, actual[userId6])
			assert.Equal(t, 0, actual[userId7])
		})
	})
}
