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
					UserId:         userId,
					SubmissionType: SubmissionType{Coefficient: 0.8},
					Score:          1000000,
				},
				{
					UserId:         userId,
					SubmissionType: SubmissionType{Coefficient: 0.9},
					Score:          980000,
				},
				{
					UserId:         userId,
					SubmissionType: SubmissionType{Coefficient: 1.0},
					Score:          950000,
				},
			},
		}

		expect := map[user.UserId]Submission{
			userId: {
				UserId: userId,
				SubmissionType: SubmissionType{
					Coefficient: 1.0,
				},
				Score: 950000,
			},
		}

		actual := challenge.AdoptedSubmissions()

		assert.Equal(t, expect, actual)
	})
}

func Test(t *testing.T) {
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
					UserId:         userId1,
					SubmissionType: SubmissionType{Coefficient: 1.0},
					Score:          500,
				},
				{
					UserId:         userId2,
					SubmissionType: SubmissionType{Coefficient: 1.0},
					Score:          400,
				},
				{
					UserId:         userId3,
					SubmissionType: SubmissionType{Coefficient: 1.0},
					Score:          300,
				},
				{
					UserId:         userId4,
					SubmissionType: SubmissionType{Coefficient: 1.0},
					Score:          300,
				},
				{
					UserId:         userId5,
					SubmissionType: SubmissionType{Coefficient: 1.0},
					Score:          200,
				},
				{
					UserId:         userId6,
					SubmissionType: SubmissionType{Coefficient: 1.0},
					Score:          100,
				},
				{
					UserId:         userId7,
					SubmissionType: SubmissionType{Coefficient: 1.0},
					Score:          0,
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
