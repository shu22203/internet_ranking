package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func timeMustParse(value string) time.Time {
	time, err := time.Parse("2006/01/02 15:04:05", value)

	if err != nil {
		panic(err)
	}

	return time
}

func TestFirstStartAt(t *testing.T) {
	t.Run("部門が1つ以上存在するとき", func(t *testing.T) {
		irs := IrSections{
			sections: []Section{
				{startAt: timeMustParse("2020/04/01 12:00:00")},
				{startAt: timeMustParse("2020/04/01 10:00:00")},
				{startAt: timeMustParse("2020/04/02 10:00:00")},
			},
		}
		actual, err := irs.FirstStartAt()

		t.Run("全部門通しての開始時刻を返す", func(t *testing.T) {
			assert.Equal(t, timeMustParse("2020/04/01 10:00:00"), actual)
		})

		t.Run("errorは返さない", func(t *testing.T) {
			assert.Nil(t, err)
		})
	})

	t.Run("部門が存在しないとき", func(t *testing.T) {
		irs := IrSections{
			sections: []Section{},
		}
		_, err := irs.FirstStartAt()

		t.Run("errorを返す", func(t *testing.T) {
			assert.NotNil(t, err)
		})
	})
}

func TestLastEndAt(t *testing.T) {
	t.Run("部門が1つ以上存在するとき", func(t *testing.T) {
		irs := IrSections{
			sections: []Section{
				{endAt: timeMustParse("2020/04/01 12:00:00")},
				{endAt: timeMustParse("2020/04/01 10:00:00")},
				{endAt: timeMustParse("2020/04/02 10:00:00")},
			},
		}
		actual, err := irs.LastEndAt()

		t.Run("全部門通しての終了時刻を返す", func(t *testing.T) {
			assert.Equal(t, timeMustParse("2020/04/02 10:00:00"), actual)
		})

		t.Run("errorは返さない", func(t *testing.T) {
			assert.Nil(t, err)
		})
	})

	t.Run("部門が存在しないとき", func(t *testing.T) {
		irs := IrSections{
			sections: []Section{},
		}
		_, err := irs.LastEndAt()

		t.Run("errorを返す", func(t *testing.T) {
			assert.NotNil(t, err)
		})
	})
}
