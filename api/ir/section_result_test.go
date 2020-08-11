package ir

import (
	"github.com/shu22203/internet_ranking/challenge"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTotalPoint(t *testing.T) {
	t.Run("部門個人結果の合計点を返す", func(t *testing.T) {
		isr := SectionIndividualResult{
			challenge.NewChallengeId(): 12,
			challenge.NewChallengeId(): 15,
			challenge.NewChallengeId(): 10,
		}

		assert.Equal(t, 37, isr.TotalPoint())
	})
}
