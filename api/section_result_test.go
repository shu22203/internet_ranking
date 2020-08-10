package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTotalPoint(t *testing.T) {
	t.Run("部門個人結果の合計点を返す", func(t *testing.T) {
		isr := SectionIndividualResult{
			NewChallengeId(): 12,
			NewChallengeId(): 15,
			NewChallengeId(): 10,
		}

		assert.Equal(t, 37, isr.TotalPoint())
	})
}
